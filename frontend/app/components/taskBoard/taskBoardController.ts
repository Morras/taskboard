///<reference path="../../references.ts" />

module tasks.controllers {

    export class TaskBoardController {

        scope;
        timeout: ng.ITimeoutService;
        taskBoardFactory: tasks.factories.ITaskBoardFactory;
        
        tasks: tasks.shared.Task[];
        lastPeriodsTasks: tasks.shared.Task[]

        jubilation: string;
        jubilations:string[] = [
            'WOOHOO',
            'Way to go',
            'I rock',
            'Keep up the good work',
            'YAY',
            'Well done',
            'I am the king of the world',
            'And I saw that it was good',
            'Jubilation!'
        ];

        constructor($scope, $timeout, taskBoardFactory) {
            this.scope = $scope;
            this.scope.vm = this;
            this.taskBoardFactory = taskBoardFactory;
            this.timeout = $timeout;

            this.jubilation = this.jubilations[Math.floor(Math.random() * 100) % this.jubilations.length];
            this.tick();
        }

        addWork(task: tasks.shared.Task, delta) {
            task.workDone += delta;
        };

        tick() {
            this.tasks = this.taskBoardFactory.getTasksInCurrentPeriod();
            this.lastPeriodsTasks = this.taskBoardFactory.getTasksInPreviousPeriod();
            this.timeout(() => this.tick(), 1000);
        }
    }
}
tasks.taskBoardApp.controller(
    'taskBoardController', 
    [
        '$scope', 
        '$timeout', 
        'taskBoardFactory',
        ($scope, $timeout, TaskBoardFactory) => new tasks.controllers.TaskBoardController($scope, $timeout, TaskBoardFactory)
    ]
);