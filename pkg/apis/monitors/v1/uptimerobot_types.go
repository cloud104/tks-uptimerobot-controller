/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UptimeRobotFinalizer ...
const UptimeRobotFinalizer = "uptimerobot.k8s.io"

// UptimeRobotSpec defines the desired state of UptimeRobot
type UptimeRobotSpec struct {
	StatusPage    UptimeStatusPage   `json:"statusPage"`
	Hosts         []UptimeRobotHosts `json:"hosts"`
	AlertContacts []AlertContacts    `json:"alertContact"`
}

// UptimeStatusPage ...
type UptimeStatusPage struct {
	URL          string `json:"url"`
	FriendlyName string `json:"friendlyName"`
}

// UptimeRobotHosts ...
type UptimeRobotHosts struct {
	URL          string `json:"url"`
	FriendlyName string `json:"friendlyName"`
	Type         string `json:"type,omitempty"`
}

// AlertContact ...
type AlertContacts struct {
	ID         string `json:"id"`
	Threshold  string `json:"threshold"`
	Recurrence string `json:"recurrence"`
}

// UptimeRobotStatus defines the observed state of UptimeRobot
type UptimeRobotStatus struct {
	Name string `json:"name,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UptimeRobot is the Schema for the uptimerobots API
// +k8s:openapi-gen=true
type UptimeRobot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UptimeRobotSpec   `json:"spec,omitempty"`
	Status UptimeRobotStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UptimeRobotList contains a list of UptimeRobot
type UptimeRobotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UptimeRobot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UptimeRobot{}, &UptimeRobotList{})
}
