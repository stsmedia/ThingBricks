'use strict';

/**
 * @ngdoc service
 * @name thingbricks.AuthService
 * @description
 * # AuthService
 * Factory in thingbricks.
 */
angular.module('thingbricks')
    .factory('authService', function ($http, ENV, $q, Session) {
        var authorizationResult = false;

        return {
            initialize: function () {
                OAuth.initialize('26eh_bICYcNNmNKBHxlbXsXsdGs', {cache: true});
                authorizationResult = OAuth.create('google');
            },
            isAuthenticated: function () {
                return !!Session.id;
            },
            isAuthorized: function (authorizedRoles) {
                if (!angular.isArray(authorizedRoles)) {
                    authorizedRoles = [authorizedRoles];
                }
                return (oauth.isAuthenticated() &&
                    authorizedRoles.indexOf(Session.userRole) !== -1);
            },
            authenticate: function (type, service) {
                var deferred = $q.defer();
                var promise = OAuth.popup(service, {cache:true});
                promise.done( function (result) {
                    $http.post(ENV.apiEndpoint + "/" + type, {network: service, accessToken: result.access_token})
                        .then(function (res) {
                            var login = res.data;
                            Session.create(login.id, login.firstName, login.lastName, login.picture);
                            deferred.resolve(login);
                        })
                }).fail( function(error) {
                    Session.destroy();
                    deferred.reject(error);
                });
                return deferred.promise;
            },
            logout: function () {
                OAuth.clearCache('google');
                authorizationResult = false;
                Session.destroy();
            }
        }
    })
    .service('Session', function () {
        this.create = function (id, firstName, lastName, picture, userRole) {
            this.id = id;
            this.firstName = firstName;
            this.lastName = lastName;
            this.picture = picture;
            this.userRole = userRole;
        };
        this.destroy = function () {
            this.id = null;
            this.firstName = null;
            this.lastName = null;
            this.picture = null;
            this.userRole = null;
        };
        return this;
    })
