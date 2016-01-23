///<references< path="../../references.ts" />

module tasks.controllers{

    import factories = tasks.factories;

    export class ManageRecurringTasksController{

        //Sets to any as we are extending the scope
        scope: any;
        taskBoardFactory: factories.ITaskBoardFactory;

        constructor($scope: ng.IScope, taskBoardFactory: factories.ITaskBoardFactory){

            this.scope = $scope;
            this.taskBoardFactory = taskBoardFactory;

            console.log(this.scope);
            console.log(this.scope.id);
        }
    }
}

tasks.taskBoardApp.controller(
    "manageRecurringTasksController", 
    [
        '$scope', 
        'taskBoardFactory',
        ($scope, TaskBoardFactory) => new tasks.controllers.ManageRecurringTasksController($scope, TaskBoardFactory)
    ]
);
