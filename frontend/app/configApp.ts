///<reference path="references.ts" />


module tasks.config{

    export var taskBoardApp = angular.module('taskBoardApp', ['ngRoute', 'ui.bootstrap']);

    /**
     * as controllers, factories and directives are all defined in their
     * own files, consider also adding them to taskBoardApp in the same
     * file, just after the modules.
     * This would keep bootstrapping out of the configuration that should
     * be done here.
     */

    //Add controllers
    taskBoardApp.controller("manageRecurringTasksController", ['$scope', 'taskBoardFactory',
        ($scope, TaskBoardFactory) => new tasks.controllers.ManageRecurringTasksController($scope, TaskBoardFactory)]);
    taskBoardApp.controller('editTaskController', ['$scope', 'taskBoardFactory', '$routeParams',
        ($scope, TaskBoardFactory, $routeParams) => new tasks.controllers.EditTaskController($scope, TaskBoardFactory, $routeParams)]);
    taskBoardApp.controller('taskBoardController', ['$scope', '$timeout', 'taskBoardFactory',
        ($scope, $timeout, TaskBoardFactory) => new tasks.controllers.TaskBoardController($scope, $timeout, TaskBoardFactory)]);

    //Add factories
    taskBoardApp.factory("taskBoardFactory", [
        () => new tasks.factories.TaskBoardFactory()]);
    taskBoardApp.factory("addWorkModalFactory", ['$modal',
        ($modal) => new tasks.factories.AddWorkModalFactory($modal)]);

    //Add directives
    taskBoardApp.directive('taskDirective', ['addWorkModalFactory',
        (AddWorkModalFactory) => new tasks.directives.TaskDirective(AddWorkModalFactory)]);

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

