angular.module('Todo.auth', [])

.factory('Auth', function($http, $state){
  return {
    login: function(obj){
      $http({
        method: 'POST',
        url: '/auth/login',
        data: obj
      }).success(function(resp){
        console.log('I am the resp', resp);
      })
    },
    signup: function(obj){
      $http({
        method: 'POST',
        url: '/auth/signup',
        data: obj
      }).then(function(resp){
        console.log('I am the resp', resp);
      })
    },
    logout: function(){
      $state.go('login');
    },
    getData: function(){

    }
  }
})
