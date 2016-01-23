///<reference path="references.ts" />


module tasks{

    export var taskBoardApp = angular.module('taskBoardApp', ['ngRoute', 'ui.bootstrap']);

    taskBoardApp.config(['$routeProvider', function ($routeProvider: ng.route.IRouteProvider) {
        $routeProvider
            .when('/',
            {
                controller: 'taskBoardController',
                templateUrl: '/app/components/taskBoard/taskBoardView.html'
            })
            .when('/editTask/:id',
            {
                controller: 'editTaskController',
                templateUrl: '/app/components/editTask/editTaskView.html'
            })
            .when('/manageRecurringTasks',
            {
                controller: 'manageRecurringTasksController',
                templateUrl: '/app/components/manageRecurringTasks/manageRecurringTasksView.html'
            })
            .otherwise({ redirectTo: '/' });
    }]);
}

