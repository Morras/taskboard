///<reference path="../../references.ts" />

module tasks.directives{

    export class TaskDirective{

        constructor(AddWorkModalFactory: tasks.factories.AddWorkModalFactory){
            console.log("Constructing directive");
            return this.createDirective(AddWorkModalFactory);
        }

        private createDirective(AddWorkModalFactory: tasks.factories.AddWorkModalFactory): any{
            console.log("Creating directive");
            return {

                scope: {
                    task: '='
                },

                controller: function($scope, $element){
                    $scope.addWorkModal =  function () {
                        AddWorkModalFactory.showModal().then(function (result) {
                            var delta = result.delta;
                            if (!isNaN(delta)) {
                                $scope.task.progress = parseInt($scope.task.progress) + parseInt(delta);
                            }
                        });
                    }},

                templateUrl: 'app/components/taskBoard/taskDirectiveTemplate.html'

            };
        }
    }
}
