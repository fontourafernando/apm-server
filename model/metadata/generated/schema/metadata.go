// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package schema

const ModelSchema = `{
    "$id": "doc/spec/metadata.json",
    "title": "Metadata",
    "description": "Metadata concerning the other objects in the stream.",
    "type": ["object"],
    "properties": {
        "service": {
                "$id": "doc/spec/service.json",
    "title": "Service",
    "type": ["object", "null"],
    "properties": {
        "agent": {
            "description": "Name and version of the Elastic APM agent",
            "type": ["object", "null"],
            "properties": {
                "name": {
                    "description": "Name of the Elastic APM agent, e.g. \"Python\"",
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "version": {
                    "description": "Version of the Elastic APM agent, e.g.\"1.0.0\"",
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "ephemeral_id": {
                    "description": "Free format ID used for metrics correlation by some agents",
                    "type": ["string", "null"],
                    "maxLength": 1024
                }
            }
        },
        "framework": {
            "description": "Name and version of the web framework used",
            "type": ["object", "null"],
            "properties": {
                "name": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "version": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                }
            }
        },
        "language": {
            "description": "Name and version of the programming language used",
            "type": ["object", "null"],
            "properties": {
                "name": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "version": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                }
            }
        },
        "name": {
            "description": "Immutable name of the service emitting this event",
            "type": ["string", "null"],
            "pattern": "^[a-zA-Z0-9 _-]+$",
            "maxLength": 1024
        },
        "environment": {
            "description": "Environment name of the service, e.g. \"production\" or \"staging\"",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "runtime": {
            "description": "Name and version of the language runtime running this service",
            "type": ["object", "null"],
            "properties": {
                "name": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "version": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                }
            }
        },
        "version": {
            "description": "Version of the service emitting this event",
            "type": ["string", "null"],
            "maxLength": 1024
        }
    },
            "type": "object",
            "required": ["name", "agent"],
            "properties.name.type": "string",
            "properties.agent.type": "string",
            "properties.agent.required": ["name", "version"],
            "properties.agent.properties.name.type": "string",
            "properties.agent.properties.version.type": "string",
            "properties.runtime.required": ["name", "version"],
            "properties.runtime.properties.name.type": "string",
            "properties.runtime.properties.version.type": "string",
            "properties.language.required": ["name"],
            "properties.language.properties.name.type": "string"
        },
        "process": {
              "$id": "doc/spec/process.json",
  "title": "Process",
  "type": ["object", "null"],
  "properties": {
      "pid": {
          "description": "Process ID of the service",
          "type": ["integer"]
      },
      "ppid": {
          "description": "Parent process ID of the service",
          "type": ["integer", "null"]
      },
      "title": {
          "type": ["string", "null"],
          "maxLength": 1024
      },
      "argv": {
        "description": "Command line arguments used to start this process",
        "type": ["array", "null"],
        "minItems": 0,
        "items": {
           "type": "string"
        }
    }
  },
  "required": ["pid"]
        },
        "system": {
                "$id": "doc/spec/system.json",
    "title": "System",
    "type": ["object", "null"],
    "properties": {
        "architecture": {
            "description": "Architecture of the system the agent is running on.",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "hostname": {
            "description": "Hostname of the system the agent is running on. Will be ignored if kubernetes information is set.",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "name": {
            "description": "Name of the system the agent is running on. Will be set to hostname or derived from kubernetes information if not provided.",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "platform": {
            "description": "Name of the system platform the agent is running on.",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "container": {
            "properties": {
                "id" : {
                    "description": "Container ID",
                    "type": ["string"],
                    "maxLength": 1024
                }
            },
            "required": ["id"]
        },
        "kubernetes": {
            "properties": {
                "namespace": {
                    "description": "Kubernetes namespace",
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "pod":{
                    "properties": {
                        "name": {
                            "description": "Kubernetes pod name",
                            "type": ["string", "null"],
                            "maxLength": 1024
                        },
                        "uid": {
                            "description": "Kubernetes pod uid",
                            "type": ["string", "null"],
                            "maxLength": 1024
                        }
                    }
                },
                "node":{
                    "properties": {
                        "name": {
                            "description": "Kubernetes node name",
                            "type": ["string", "null"],
                            "maxLength": 1024
                        }
                    }
                }
            }
        }
    }
        },
        "user": {
            "description": "Describes the authenticated User for a request.",
                "$id": "docs/spec/user.json",
    "title": "User",
    "type": ["object", "null"],
    "properties": {
        "id": {
            "description": "Identifier of the logged in user, e.g. the primary key of the user",
            "type": ["string", "integer", "null"],
            "maxLength": 1024
        },
        "email": {
            "description": "Email of the logged in user",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "username": {
            "description": "The username of the logged in user",
            "type": ["string", "null"],
            "maxLength": 1024
        }
    }
        },
        "labels": {
                "$id": "doc/spec/tags.json",
    "title": "Tags",
    "type": ["object", "null"],
    "description": "A flat mapping of user-defined tags with string, boolean or number values.",
    "patternProperties": {
        "^[^.*\"]*$": {
            "type": ["string", "boolean", "number", "null"],
            "maxLength": 1024
        }
    },
    "additionalProperties": false
        }
    },
    "required": ["service"]
}
`
