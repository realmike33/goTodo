angular.module('Todo.login', [])

.controller('LoginCtrl', function($scope, Auth){
  $scope.login = function(){
    Auth.login({username: 'Yeah', password: 'yup', todo: 'yeah'})
  }
})
