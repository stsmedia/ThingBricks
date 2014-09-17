'use strict';

/**
 * @ngdoc service
 * @name thingbricks.ProjectService
 * @description
 * # ProjectService
 * Factory in thingbricks.
 */
angular.module('thingbricks')
    .factory('projectService', function ($http, ENV, $q) {
        return {
            listAll: function (accountId) {
                var promise = $http.get(ENV.apiEndpoint + '/accounts/' + accountId + '/projects')
                    .success(function(res) {
                        return res;
                    })
                    .error(function(err) {
                        console.log(err);
                        return err;
                    });
                return promise;
            },
            add: function(project) {
                if (!project.id) {
                    var promise = $http.post(ENV.apiEndpoint + '/accounts/' + project.accountId + '/projects', project)
                        .success(function (res) {
                            return res;
                        }).error(function (err) {
                            console.log(err);
                            return err;
                        });
                    return promise;
                } else {
                    var promise = $http.put(ENV.apiEndpoint + '/accounts/' + project.accountId + '/projects/' + project.id, project)
                        .success(function (res) {
                            return res;
                        }).error(function(err) {
                            console.log(err);
                            return err;
                        });
                    return promise;
                }
            },
            delete: function(project) {
                var promise = $http.delete(ENV.apiEndpoint + '/accounts/' + project.accountId + '/projects/' + project.id, project)
                    .success(function (res) {
                        return res;
                    }).error(function(err) {
                        console.log(err);
                        return err;
                    });
                return promise;
            }
        }
    })
