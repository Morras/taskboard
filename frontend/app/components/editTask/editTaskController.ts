///<references< path="../../references.ts" />

module tasks.controllers{

    import factories = tasks.factories;

    export class EditTaskController
    {
        scope: any;
        taskBoardFactor: factories.ITaskBoardFactory;
        routeParams: any;

        constructor($scope: ng.IScope, taskBoardFactory: tasks.factories.ITaskBoardFactory, $routeParams: ng.route.IRouteParamsService){
            this.scope = $scope;
            this.taskBoardFactor = taskBoardFactory;
            this.routeParams = $routeParams;
            this.scope.taskId = this.routeParams.id;

            console.log(this.routeParams);
            console.log(this.routeParams.id);
        }
    }
}

tasks.taskBoardApp.controller(
    'editTaskController', 
    [
        '$scope', 
        'taskBoardFactory', 
        '$routeParams',
        ($scope, TaskBoardFactory, $routeParams) => new tasks.controllers.EditTaskController($scope, TaskBoardFactory, $routeParams)
    ]
);