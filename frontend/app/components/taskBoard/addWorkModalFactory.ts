///<reference path="../../references.ts" />


module tasks.factories{

    export interface IAddWorkModalFactory{
        showModal(customModalDefaults, customModalOptions);
    }

    export class AddWorkModalFactory implements IAddWorkModalFactory{
        private modalDefaults = {
            backdrop: 'static',
            keyboard: true,
            modalFade: true,
            templateUrl: '/app/components/taskBoard/addWorkModalTemplate.html'
        };

        private modalOptions = {
            cancelButtonText: 'cancel',
            okButtonText: 'OK',
            headerText: 'Add work to task',
            bodyText: 'How much work has been done since last update?'
        };

        private modal;

        constructor($modal){
            this.modal = $modal;
        }

        showModal = function () {
            //Create temporary objects as we are working with a singleton service
            var tempModalDefaults: any = {};
            var tempModalOptions: any = {};
            //Map angular-ui modal custom defaults to modal defaults defined in service
            angular.extend(tempModalDefaults, this.modalDefaults);
            //Map modal.html $scope custom properties to defaults defined in service
            angular.extend(tempModalOptions, this.modalOptions);
            if (!tempModalDefaults.controller) {
                tempModalDefaults.controller = function ($scope, $modalInstance) {
                    $scope.modalOptions = tempModalOptions;
                    $scope.modalOptions.ok = function (result) {
                        if (!result) {
                            result = {};
                        }
                        result.delta = $scope.delta;
                        $modalInstance.close(result);
                    };
                    $scope.modalOptions.close = function () {
                        $modalInstance.dismiss('cancel');
                    };
                }
            }
            return this.modal.open(tempModalDefaults).result;
        }
    }
}


tasks.taskBoardApp.factory(
    "addWorkModalFactory", 
    [
        '$modal',
        ($modal) => new tasks.factories.AddWorkModalFactory($modal)
    ]
);