///<references path="../../references.ts" />

module tasks.factories{

    export interface ITaskBoardFactory {
        getLastPeriodsTasks() : tasks.shared.Task[];
        getMustDos() : tasks.shared.Task[];
        getWantToDos() : tasks.shared.Task[];
        getTask(id: Number, isRecurring: Boolean) : tasks.shared.Task;
    }

    export class TaskBoardFactory implements ITaskBoardFactory{

        private lastPeriodsTasks: tasks.shared.Task[] = [
            {
                id: 1,
                text: 'Vask trappen 1 gang',
                progress: [{ period: new tasks.shared.Period(51, 2015), progress: 1}],
                totalWork: 1,
                type: tasks.shared.TaskType.MUST,
                recurring: true,
                firstPeriod: new tasks.shared.Period(51, 2015)
            },
            {
                id: 2,
                text: 'Læs 50 sider faglitterært',
                progress: [{ period: new tasks.shared.Period(51, 2015), progress: 52}],
                totalWork: 50,
                type: tasks.shared.TaskType.MUST,
                recurring: true,
                firstPeriod: new tasks.shared.Period(51, 2015)
            }
        ];

        private mustDos: tasks.shared.Task[] = [
            {
                id: 3,
                text: 'Vask trappen 1 gang',
                progress:  [{ period: new tasks.shared.Period(52, 2015), progress: 1}],
                totalWork: 1,
                type: tasks.shared.TaskType.MUST,
                recurring: true,
                firstPeriod: new tasks.shared.Period(51, 2015)
            }
        ];

        private wantToDos: tasks.shared.Task[] = [
            {
                id: 4,
                text: 'Læse 50 sider faglitterært',
                progress:  [{ period: new tasks.shared.Period(52, 2015), progress: 1}],
                totalWork: 50,
                type: tasks.shared.TaskType.MUST,
                recurring: true,
                firstPeriod: new tasks.shared.Period(51, 2015)
            },
            {
                id: 5,
                text: 'Træne 10 gange',
                progress:  [{ period: new tasks.shared.Period(52, 2015), progress: 1}],
                totalWork: 10,
                type: tasks.shared.TaskType.MUST,
                recurring: true,
                firstPeriod: new tasks.shared.Period(51, 2015)
            },
            {
                id: 6,
                text: 'Kode 8 timer',
                progress:  [{ period: new tasks.shared.Period(52, 2015), progress: 1}],
                totalWork: 8,
                type: tasks.shared.TaskType.MUST,
                recurring: true,
                firstPeriod: new tasks.shared.Period(51, 2015)
            }
        ];

        constructor(){
            setInterval(() => this.fetchData(), 10000); //10 seconds
        }

        private fetchData() {
            console.log("fetching")
        }

        getLastPeriodsTasks(){
            return this.lastPeriodsTasks;
        }

        getMustDos() {
            return this.mustDos;
        }

        getWantToDos(){
            return this.wantToDos;
        }

        getTask(id: Number, isRecurring: Boolean){
            if ( isRecurring ){
                //Fetch task from server
            } else {
                for ( var task in this.mustDos ){
                    if ( task.id = id ){
                        return task;
                    }
                }
                for ( task in this.wantToDos ){
                    if ( task.id = id ){
                        return task;
                    }
                }
            }
            return null;
        }

    }
}


tasks.taskBoardApp.factory(
    "taskBoardFactory", 
    [
        () => new tasks.factories.TaskBoardFactory()
    ]
);