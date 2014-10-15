'use strict';

/**
 * @ngdoc function
 * @name thingbricks.controller:DataStreamCtrl
 * @description
 * # DataStreamCtrl
 * Controller of thingbricks
 */
angular.module('thingbricks')
    .controller('DataStreamCtrl', function ($scope, dataStreamService) {
        $scope.dataStreams = [];
        $scope.selectedDataStream = null;
        $scope.message = "";
        $scope.showDataStreamForm = false;
        $scope.editDataStreamForm = false;
        $scope.allDataStreams = function () {
            dataStreamService.listAll(1)
                .then(function (result) {
                    for (var i = 0; i < result.data.length; i++) {
                        $scope.dataStreams.push(result.data[i]);
                    }
                });
        }
        $scope.addDataStream = function () {
            dataStreamService.add($scope.selectedDataStream, $scope.selectedProject)
                .then(function (result) {
                    console.log(result);
                    $scope.message = "data stream created";
                    $scope.showDataStreamForm = false;
                    if (!$scope.selectedDataStream.id)
                        $scope.dataStreams.push(result.data);
                    $scope.selectedDataStream = null;
                });
        }
        $scope.viewDataStream = function (dataStream, edit) {
            $scope.selectedDataStream = dataStream;
            $scope.showDataStreamForm = true;
            $scope.editDataStreamForm = !edit;
        }
        $scope.deleteDataStream = function (dataStream, project) {
            dataStreamService.delete(dataStream, project)
                .then(function (result) {
                    console.log(result);
                    $scope.message = "data stream deleted";
                    $scope.showDataStreamForm = false;
                    $scope.dataStreams.pop(dataStream);
                    $scope.selectedDataStream = null;
                });
        }
    })