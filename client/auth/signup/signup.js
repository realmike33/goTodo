angular.module('Todo.signup', [])

.controller('SignupCtrl', function($scope, Auth){
  $scope.signup = function(){
    Auth.signup({username: 'Mike', password: 'Yeah'});
  }
})
