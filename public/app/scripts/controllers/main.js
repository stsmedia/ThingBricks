'use strict';

/**
 * @ngdoc function
 * @name thingbricks.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of thingbricks
 */
angular.module('thingbricks')
    .controller('MainCtrl', function ($scope, USER_ROLES, $rootScope, AUTH_EVENTS, authService) {
        authService.initialize();

        $scope.currentUser = null;
        $scope.userRoles = USER_ROLES;
        $scope.isAuthorized = authService.isAuthorized;

        $scope.setCurrentUser = function (user) {
            $scope.currentUser = user;
        };
        $scope.auth = function (type, service) {
            authService.authenticate(type, service).then(function (login) {
                $rootScope.$broadcast(AUTH_EVENTS.loginSuccess);
                $scope.setCurrentUser(login);
            }, function (error) {
                console.log(error);
                $rootScope.$broadcast(AUTH_EVENTS.loginFailed);
            });
        };
        $scope.logout = function() {
            console.log('logging out')
            authService.logout();
            $scope.setCurrentUser(null);
        }
    });
