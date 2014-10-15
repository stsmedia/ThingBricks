'use strict';

/**
 * @ngdoc function
 * @name thingbricks.controller:ProjectCtrl
 * @description
 * # ProjectCtrl
 * Controller of thingbricks
 */
angular.module('thingbricks')
    .controller('ProjectCtrl', function ($scope, projectService) {
        $scope.projects = [];
        $scope.selectedProject = null;
        $scope.message = "";
        $scope.showProjectForm = false;
        $scope.editProjectForm = false;
        $scope.allProjects = function () {
            projectService.listAll(1)
                .then(function (result) {
                    for (var i = 0; i < result.data.length; i++) {
                        $scope.projects.push(result.data[i]);
                    }
                });
        }
        $scope.addProject = function () {
            projectService.add($scope.selectedProject)
                .then(function (result) {
                    console.log(result);
                    $scope.message = "project created";
                    $scope.showProjectForm = false;
                    if (!$scope.selectedProject.id)
                        $scope.projects.push(result.data);
                    $scope.selectedProject = null;
                });
        }
        $scope.viewProject = function (project, edit) {
            $scope.selectedProject = project;
            $scope.showProjectForm = true;
            $scope.editProjectForm = !edit;
        }
        $scope.deleteProject = function (project) {
            projectService.delete(project)
                .then(function (result) {
                    console.log(result);
                    $scope.message = "project deleted";
                    $scope.showProjectForm = false;
                    $scope.projects.pop(project);
                    $scope.selectedProject = null;
                });
        }
    });
