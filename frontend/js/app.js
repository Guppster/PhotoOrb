var main = angular.module('myApp', ['ngRoute', 'three60Controller']);

main.config(['$routeProvider',
    function ($routeProvider) {
        $routeProvider.when('/home', {
            templateUrl: 'views/home.html',
            controller: 'Three60Controller'
        }).when('/preview', {
            templateUrl: 'views/preview.html',
            controller: 'Three60Controller'
        }).otherwise({
            redirectTo: '/home'
        });
    }]);
main.run(['$rootScope', '$location',
    function ($rootScope, $location) {
        $rootScope.$on('$routeChangeStart', function (event, currRoute, prevRoute) {

        });
    }]);
