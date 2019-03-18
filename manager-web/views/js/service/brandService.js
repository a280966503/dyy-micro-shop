// 定义服务层:
app.service("brandService",function($http){
	this.findAll = function(){
		return $http.get("../login/findAll.do");
	}
	
	this.findByPage = function(page,rows){
		return $http.get("../login/findByPage.do?page="+page+"&rows="+rows);
	}
	
	this.save = function(entity){
		return $http.post("../login/save.do",entity);
	}
	
	this.update=function(entity){
		return $http.post("../login/update.do",entity);
	}
	
	this.findById=function(id){
		return $http.get("../login/findById.do?id="+id);
	}
	
	this.dele = function(ids){
		return $http.get("../login/delete.do?ids="+ids);
	}
	
	this.search = function(page,rows,searchEntity){
		return $http.post("../login/search.do?page="+page+"&rows="+rows,searchEntity);
	}
	
	this.selectOptionList = function(){
		return $http.get("../login/selectOptionList.do");
	}
});