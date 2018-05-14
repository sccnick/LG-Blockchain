//SPDX-License-Identifier: Apache-2.0

var lg = require('./controller.js');

module.exports = function(app){

  app.get('/get_lg/:id', function(req, res){
    lg.get_lg(req, res);
  });
  app.get('/add_lg/:lg', function(req, res){
    lg.add_lg(req, res);
  });
  app.get('/get_all_lg', function(req, res){
    lg.get_all_lg(req, res);
  });
  app.post('/add_lgpost',(req,res)=>{
    console.log(req)
    lg.add_lgpost(req,res)
  })
  // app.get('/change_holder/:holder', function(req, res){
  //   lg.change_holder(req, res);
  // });
}
