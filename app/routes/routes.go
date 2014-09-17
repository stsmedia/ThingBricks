// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tLoginController struct {}
var LoginController tLoginController


func (_ tLoginController) Signin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("LoginController.Signin", args).Url
}

func (_ tLoginController) FindByAccountGroup(
		accountGroupId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountGroupId", accountGroupId)
	return revel.MainRouter.Reverse("LoginController.FindByAccountGroup", args).Url
}

func (_ tLoginController) Delete(
		accountGroupId int64,
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountGroupId", accountGroupId)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("LoginController.Delete", args).Url
}


type tProjectController struct {}
var ProjectController tProjectController


func (_ tProjectController) FindByAccount(
		accountId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	return revel.MainRouter.Reverse("ProjectController.FindByAccount", args).Url
}

func (_ tProjectController) One(
		accountId int64,
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("ProjectController.One", args).Url
}

func (_ tProjectController) Add(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ProjectController.Add", args).Url
}

func (_ tProjectController) Update(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("ProjectController.Update", args).Url
}

func (_ tProjectController) Delete(
		accountId int64,
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("ProjectController.Delete", args).Url
}


type tAccountController struct {}
var AccountController tAccountController


func (_ tAccountController) All(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AccountController.All", args).Url
}

func (_ tAccountController) One(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AccountController.One", args).Url
}

func (_ tAccountController) Add(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AccountController.Add", args).Url
}

func (_ tAccountController) Update(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AccountController.Update", args).Url
}

func (_ tAccountController) Delete(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AccountController.Delete", args).Url
}


type tAccountGroupController struct {}
var AccountGroupController tAccountGroupController


func (_ tAccountGroupController) FindByAccount(
		accountId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	return revel.MainRouter.Reverse("AccountGroupController.FindByAccount", args).Url
}

func (_ tAccountGroupController) One(
		accountId int64,
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AccountGroupController.One", args).Url
}

func (_ tAccountGroupController) Add(
		accountId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	return revel.MainRouter.Reverse("AccountGroupController.Add", args).Url
}

func (_ tAccountGroupController) Update(
		accountId int64,
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AccountGroupController.Update", args).Url
}

func (_ tAccountGroupController) Delete(
		accountId int64,
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AccountGroupController.Delete", args).Url
}


type tDataStreamController struct {}
var DataStreamController tDataStreamController


func (_ tDataStreamController) FindByProject(
		accountId int64,
		projectId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	revel.Unbind(args, "projectId", projectId)
	return revel.MainRouter.Reverse("DataStreamController.FindByProject", args).Url
}

func (_ tDataStreamController) One(
		accountId int64,
		projectId int64,
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	revel.Unbind(args, "projectId", projectId)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("DataStreamController.One", args).Url
}

func (_ tDataStreamController) Add(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("DataStreamController.Add", args).Url
}

func (_ tDataStreamController) Update(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("DataStreamController.Update", args).Url
}

func (_ tDataStreamController) Delete(
		accountId int64,
		projectId int64,
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "accountId", accountId)
	revel.Unbind(args, "projectId", projectId)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("DataStreamController.Delete", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


