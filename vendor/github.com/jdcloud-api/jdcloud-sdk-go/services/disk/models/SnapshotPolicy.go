// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type SnapshotPolicy struct {

    /* 策略id (Optional) */
    Id string `json:"id"`

    /* 策略名称 (Optional) */
    Name string `json:"name"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 策略执行间隔，单位:秒 (Optional) */
    Interval int `json:"interval"`

    /* 策略生效时间。格式`YYYY-MM-DDTHH:mm:ss+xx:xx`。如`2020-02-02T20:02:00+08:00` (Optional) */
    EffectiveTime string `json:"effectiveTime"`

    /* 策略上次执行时间。格式`YYYY-MM-DDTHH:mm:ss+xx:xx`。如`2020-02-02T20:02:00+08:00` (Optional) */
    LastTriggerTime string `json:"lastTriggerTime"`

    /* 策略下次执行时间。格式`YYYY-MM-DDTHH:mm:ss+xx:xx`。如`2020-02-02T20:02:00+08:00` (Optional) */
    NextTriggerTime string `json:"nextTriggerTime"`

    /* 快照保留时间。单位:秒。0：永久保留 (Optional) */
    SnapshotLifecycle int `json:"snapshotLifecycle"`

    /* 联系人信息 (Optional) */
    ContactInfo ContactInfo `json:"contactInfo"`

    /* 策略下次执行时间。格式`YYYY-MM-DDTHH:mm:ss+xx:xx`。如`2020-02-02T20:02:00+08:00` (Optional) */
    CreateTime string `json:"createTime"`

    /* 策略下次执行时间。格式`YYYY-MM-DDTHH:mm:ss+xx:xx`。如`2020-02-02T20:02:00+08:00` (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 策略状态。1：启用 2：禁用 (Optional) */
    Status int `json:"status"`

    /* 策略绑定的disk数量 (Optional) */
    DiskCount int `json:"diskCount"`
}