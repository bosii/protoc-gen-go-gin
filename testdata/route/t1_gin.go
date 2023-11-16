// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v0.0.1
// - protoc            v4.24.0
// source: pb/t1.proto

package route

import (
	gin "github.com/gin-gonic/gin"
)

type TestServiceHandler interface {

	// ListTests 测试内容列表查询
	ListTests(c *gin.Context, req *ListTestsRequest, resp *ListTestsResponse) gin.HandlerFunc
	// ListTest2 测试内容2
	ListTest2(c *gin.Context, req *ListTestsRequest, resp *ListTestsResponse) gin.HandlerFunc
	// ListTest3 测试内容3
	ListTest3(c *gin.Context, req *ListTestsRequest, resp *ListTestsResponse) gin.HandlerFunc
}

func NewTestService(r *gin.Engine, handler TestServiceHandler) {

	r.GET("/v1alpha1", handler.ListTests())
	r.POST("/v2alpha1", handler.ListTest2())
	r.POST("/v2alpha1/:id", handler.ListTest3())
}