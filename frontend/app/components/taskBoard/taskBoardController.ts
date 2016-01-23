///<reference path="../../references.ts" />

module tasks.controllers {

    export class TaskBoardController {

        scope;
        timeout: ng.ITimeoutService;
        taskBoardFactory: tasks.factories.ITaskBoardFactory;

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
            this.taskBoardFactory = taskBoardFactory;
            this.timeout = $timeout;

            this.setupScopeFunctions();
            this.setupScopeVariables();
            this.tick();
        }

        private setupScopeFunctions() {
            this.scope.addWork = function (task: tasks.shared.Task, delta) {
                //TODO hack to get up and running after migration0
                task.progress[0].progress = task.progress[0].progress + parseInt(delta);
            };
        }

        private setupScopeVariables() {
            this.scope.jubilation = this.jubilations[Math.floor(Math.random() * 100) % this.jubilations.length];
        }

        tick() {
            this.scope.mustDos = this.taskBoardFactory.getMustDos();
            this.scope.wantToDos = this.taskBoardFactory.getWantToDos();
            this.scope.lastPeriodsTasks = this.taskBoardFactory.getLastPeriodsTasks();
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