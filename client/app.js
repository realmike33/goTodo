angular.module('Todo', [
  'ui.router',
  'Todo.auth',
  'Todo.login',
  'Todo.signup'
  ])

.config(function($stateProvider, $urlRouterProvider){
  $urlRouterProvider.otherwise('/');
  $stateProvider
    .state('login', {
      url: '/login',
      templateUrl: 'auth/login/login.html',
      controller: 'LoginCtrl'
    })
    .state('signup', {
      url: '/signup',
      templateUrl: 'auth/signup/signup.html',
      controller: 'SignupCtrl'
    })
    .state('main', {
      url: '/',
      templateUrl: 'main/main.html',
      controller: 'MainCtrl'
    })
})

.controller('MainCtrl', function($scope){

})
