///<references path="../../references.ts" />

module tasks.factories {

    export interface ITaskBoardFactory {
        getTasksInPreviousPeriod(): tasks.shared.Task[];
        getTasksInCurrentPeriod(): tasks.shared.Task[];
        jumpToPreviousPeriod();
        jumpToNextPeriod();
    }
    
    function pad(input: number) {
        var str = '' + input;
        if (str.length != 2) {
            return '0' + str;
        }
        return str
    }

    export class TaskBoardFactory implements ITaskBoardFactory {

        private date: Date = new Date();

        private previousPeriodsTasks: tasks.shared.Task[] = [];
        private tasks: tasks.shared.Task[] = [];

        private millisecondsInADay = 86400000;
        private daysInPeriod = 14;
        private millisecondsInPeriod = this.daysInPeriod * this.millisecondsInADay;
        private previousPeriodDate: Date = this.getPreviousPeriodDate();

        private http: ng.IHttpService;

        constructor($http: ng.IHttpService) {
            console.log("constructing factory")
            this.http = $http;
            this.fetchTasks()
            setInterval(() => this.fetchTasks(), 120000); //2 minutes
        }

        private getPreviousPeriodDate(): Date {
            var deltaMilliseconds = this.date.getTime() - this.millisecondsInPeriod;
            //This will glitch around DST but I will live with the one hour that might be a problem
            var previousDate = new Date(deltaMilliseconds)
            return previousDate
        }

        private fetchTasks() {
            this.fetchTasksByDateInto(this.date, this.tasks)
            this.fetchTasksByDateInto(this.previousPeriodDate, this.previousPeriodsTasks)
        }

        private fetchTasksByDateInto(date: Date, holder: tasks.shared.Task[]) {
            this.http.get(this.createRequestEndpoint(date)).
                success(function(data: tasks.shared.Task[]) {
                    console.log(data)
                    holder = data
                })
                .error(function(error) {
                    console.error(error)
                })
        }

        private createRequestEndpoint(date: Date): string {
            var year = date.getFullYear();
            var month = date.getMonth();
            var day = date.getDate();
            var datePart = year + "-" + pad(month) + "-" + pad(day)
            return '/api/task?dateInPeriod=' + datePart;
        }

        private setDateInPeriod(date: Date) {
            this.date = date;
            this.previousPeriodDate = this.getPreviousPeriodDate();
            this.fetchTasks()            
        }
        
        jumpToNextPeriod() {
            var newDate = new Date(this.date.getTime() + this.millisecondsInPeriod);
            this. setDateInPeriod(newDate);
        }
        
        jumpToPreviousPeriod() {
            var newDate = new Date(this.date.getTime() - this.millisecondsInPeriod);
            this. setDateInPeriod(newDate);
        }

        getTasksInCurrentPeriod(): tasks.shared.Task[] {
            return this.tasks;
        }

        getTasksInPreviousPeriod() {
            return this.previousPeriodsTasks;
        }
    }
}


tasks.taskBoardApp.factory(
    "taskBoardFactory",
    [
        '$http',
        ($http) => new tasks.factories.TaskBoardFactory($http)
    ]
);