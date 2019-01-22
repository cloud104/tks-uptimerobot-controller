package main

import (
	log "github.com/Sirupsen/logrus"
	api_extensions_v1beta1 "k8s.io/api/extensions/v1beta1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}) {
	log.Info("TestHandler.ObjectCreated")
	// assert the type to a Ingress object to pull out relevant data
	ingress := obj.(*api_extensions_v1beta1.Ingress)
	log.Infof("    ResourceVersion: %s", ingress.ObjectMeta.ResourceVersion)
	log.Infof("    Name: %s", ingress.ObjectMeta.Name)
	log.Infof("    Namespace: %s", ingress.ObjectMeta.Namespace)
	log.Infof("    Host: %s", ingress.Spec.Rules[0].Host)
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}
