'use strict';

/**
 * @ngdoc service
 * @name thingbricks.DataStreamService
 * @description
 * # DataStreamService
 * Factory in thingbricks.
 */
angular.module('thingbricks')
    .factory('dataStreamService', function ($http, ENV, $q) {
        return {
            listAll: function (accountId, projectId) {
                var promise = $http.get(ENV.apiEndpoint + '/accounts/' + accountId + '/projects/' + projectId + '/datastreams')
                    .success(function(res) {
                        return res;
                    })
                    .error(function(err) {
                        console.log(err);
                        return err;
                    });
                return promise;
            },
            add: function(dataStream) {
                if (!dataStream.id) {
                    var promise = $http.post(ENV.apiEndpoint + '/accounts/' + project.accountId + '/projects/' + projectId + '/datastreams', dataStream)
                        .success(function (res) {
                            return res;
                        }).error(function (err) {
                            console.log(err);
                            return err;
                        });
                    return promise;
                } else {
                    var promise = $http.put(ENV.apiEndpoint + '/accounts/' + project.accountId + '/projects/' + projectId + '/datastreams/' + dataStream.id, dataStream)
                        .success(function (res) {
                            return res;
                        }).error(function(err) {
                            console.log(err);
                            return err;
                        });
                    return promise;
                }
            },
            delete: function(dataStream) {
                var promise = $http.delete(ENV.apiEndpoint + '/accounts/' + project.accountId + '/projects/' + projectId + '/datastreams/' + dataStream.id, dataStream)
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
