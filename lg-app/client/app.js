// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.queryAllLG = function(){

		appFactory.queryAllLG(function(data){

			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
		
			$scope.all_lg = array;
		});
	}

	$scope.queryLG = function(){

		var id = $scope.lg_id;

		appFactory.queryLG(id, function(data){
			$scope.query_lg = data;

			if ($scope.query_lg == "Could not locate lg"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordLG = function(){

		appFactory.recordLG($scope.lg, function(data){
			$scope.create_lg = data;
			$("#success_create").show();
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllLG = function(callback){

    	$http.get('/get_all_lg/').success(function(output){
			callback(output)
		});
	}

	factory.queryLG = function(id, callback){
    	$http.get('/get_lg/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordLG = function(data, callback){
		var lg = data.id + "-" + data.amount + "-" + data.dealer + "-" + data.beneficial + "-" + data.name;

		var body = {
			key:data.id,
			dealer_names:data.dealer,
			beneficial_names:data.beneficial,
			guarantee_amount:data.amount,
			document_name:data.name
		}
		var config = {
			headers:{
				'Content-Type':'application/json'
			}
		}
		var req = {
			method: 'POST',
			url: '/add_lgpost',
			headers: {
			  'Content-Type': 'application/json'
			},
			data: {
				key:data.id,
				dealer_names:data.dealer,
				beneficial_names:data.beneficial,
				guarantee_amount:data.amount,
				document_name:data.name
			}
		   }
		// $http(req).success(function(output){
		// 	callback(output)
		// })
		
    	$http.get('/add_lg/'+lg).success(function(output){
			callback(output)
		});
	}

	return factory;
});


