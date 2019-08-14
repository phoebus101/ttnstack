// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	time "time"

	types "github.com/gogo/protobuf/types"
)

func (dst *PeerInfo) SetFields(src *PeerInfo, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "grpc_port":
			if len(subs) > 0 {
				return fmt.Errorf("'grpc_port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.GRPCPort = src.GRPCPort
			} else {
				var zero uint32
				dst.GRPCPort = zero
			}
		case "tls":
			if len(subs) > 0 {
				return fmt.Errorf("'tls' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TLS = src.TLS
			} else {
				var zero bool
				dst.TLS = zero
			}
		case "roles":
			if len(subs) > 0 {
				return fmt.Errorf("'roles' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Roles = src.Roles
			} else {
				dst.Roles = nil
			}
		case "tags":
			if len(subs) > 0 {
				return fmt.Errorf("'tags' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Tags = src.Tags
			} else {
				dst.Tags = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Cluster) SetFields(src *Cluster, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				newDst := &dst.ClusterIdentifiers
				var newSrc *ClusterIdentifiers
				if src != nil {
					newSrc = &src.ClusterIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClusterIdentifiers = src.ClusterIdentifiers
				} else {
					var zero ClusterIdentifiers
					dst.ClusterIdentifiers = zero
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "description":
			if len(subs) > 0 {
				return fmt.Errorf("'description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Description = src.Description
			} else {
				var zero string
				dst.Description = zero
			}
		case "attributes":
			if len(subs) > 0 {
				return fmt.Errorf("'attributes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Attributes = src.Attributes
			} else {
				dst.Attributes = nil
			}
		case "contact_info":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactInfo = src.ContactInfo
			} else {
				dst.ContactInfo = nil
			}
		case "addresses":
			if len(subs) > 0 {
				return fmt.Errorf("'addresses' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Addresses = src.Addresses
			} else {
				dst.Addresses = nil
			}
		case "secret":
			if len(subs) > 0 {
				return fmt.Errorf("'secret' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Secret = src.Secret
			} else {
				var zero string
				dst.Secret = zero
			}
		case "location":
			if len(subs) > 0 {
				newDst := &dst.Location
				var newSrc *Location
				if src != nil {
					newSrc = &src.Location
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Location = src.Location
				} else {
					var zero Location
					dst.Location = zero
				}
			}
		case "location_description":
			if len(subs) > 0 {
				return fmt.Errorf("'location_description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LocationDescription = src.LocationDescription
			} else {
				var zero string
				dst.LocationDescription = zero
			}
		case "roles":
			if len(subs) > 0 {
				return fmt.Errorf("'roles' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Roles = src.Roles
			} else {
				dst.Roles = nil
			}
		case "endpoints":
			if len(subs) > 0 {
				return fmt.Errorf("'endpoints' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Endpoints = src.Endpoints
			} else {
				dst.Endpoints = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Clusters) SetFields(src *Clusters, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "clusters":
			if len(subs) > 0 {
				return fmt.Errorf("'clusters' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Clusters = src.Clusters
			} else {
				dst.Clusters = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetClusterRequest) SetFields(src *GetClusterRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "cluster_ids":
			if len(subs) > 0 {
				newDst := &dst.ClusterIdentifiers
				var newSrc *ClusterIdentifiers
				if src != nil {
					newSrc = &src.ClusterIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClusterIdentifiers = src.ClusterIdentifiers
				} else {
					var zero ClusterIdentifiers
					dst.ClusterIdentifiers = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetClusterIdentifiersForAddressRequest) SetFields(src *GetClusterIdentifiersForAddressRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "address":
			if len(subs) > 0 {
				return fmt.Errorf("'address' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Address = src.Address
			} else {
				var zero string
				dst.Address = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListClustersRequest) SetFields(src *ListClustersRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "collaborator":
			if len(subs) > 0 {
				newDst := dst.Collaborator
				if newDst == nil {
					newDst = &OrganizationOrUserIdentifiers{}
					dst.Collaborator = newDst
				}
				var newSrc *OrganizationOrUserIdentifiers
				if src != nil {
					newSrc = src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					dst.Collaborator = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}
		case "order":
			if len(subs) > 0 {
				return fmt.Errorf("'order' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Order = src.Order
			} else {
				var zero string
				dst.Order = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateClusterRequest) SetFields(src *CreateClusterRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "cluster":
			if len(subs) > 0 {
				newDst := &dst.Cluster
				var newSrc *Cluster
				if src != nil {
					newSrc = &src.Cluster
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Cluster = src.Cluster
				} else {
					var zero Cluster
					dst.Cluster = zero
				}
			}
		case "collaborator":
			if len(subs) > 0 {
				newDst := &dst.Collaborator
				var newSrc *OrganizationOrUserIdentifiers
				if src != nil {
					newSrc = &src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					var zero OrganizationOrUserIdentifiers
					dst.Collaborator = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateClusterRequest) SetFields(src *UpdateClusterRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "cluster":
			if len(subs) > 0 {
				newDst := &dst.Cluster
				var newSrc *Cluster
				if src != nil {
					newSrc = &src.Cluster
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Cluster = src.Cluster
				} else {
					var zero Cluster
					dst.Cluster = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListClusterCollaboratorsRequest) SetFields(src *ListClusterCollaboratorsRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "cluster_ids":
			if len(subs) > 0 {
				newDst := &dst.ClusterIdentifiers
				var newSrc *ClusterIdentifiers
				if src != nil {
					newSrc = &src.ClusterIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClusterIdentifiers = src.ClusterIdentifiers
				} else {
					var zero ClusterIdentifiers
					dst.ClusterIdentifiers = zero
				}
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetClusterCollaboratorRequest) SetFields(src *GetClusterCollaboratorRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "cluster_ids":
			if len(subs) > 0 {
				newDst := &dst.ClusterIdentifiers
				var newSrc *ClusterIdentifiers
				if src != nil {
					newSrc = &src.ClusterIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClusterIdentifiers = src.ClusterIdentifiers
				} else {
					var zero ClusterIdentifiers
					dst.ClusterIdentifiers = zero
				}
			}
		case "collaborator":
			if len(subs) > 0 {
				newDst := &dst.OrganizationOrUserIdentifiers
				var newSrc *OrganizationOrUserIdentifiers
				if src != nil {
					newSrc = &src.OrganizationOrUserIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.OrganizationOrUserIdentifiers = src.OrganizationOrUserIdentifiers
				} else {
					var zero OrganizationOrUserIdentifiers
					dst.OrganizationOrUserIdentifiers = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SetClusterCollaboratorRequest) SetFields(src *SetClusterCollaboratorRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "cluster_ids":
			if len(subs) > 0 {
				newDst := &dst.ClusterIdentifiers
				var newSrc *ClusterIdentifiers
				if src != nil {
					newSrc = &src.ClusterIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ClusterIdentifiers = src.ClusterIdentifiers
				} else {
					var zero ClusterIdentifiers
					dst.ClusterIdentifiers = zero
				}
			}
		case "collaborator":
			if len(subs) > 0 {
				newDst := &dst.Collaborator
				var newSrc *Collaborator
				if src != nil {
					newSrc = &src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					var zero Collaborator
					dst.Collaborator = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Cluster_Endpoint) SetFields(src *Cluster_Endpoint, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "roles":
			if len(subs) > 0 {
				return fmt.Errorf("'roles' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Roles = src.Roles
			} else {
				dst.Roles = nil
			}

		case "endpoint":
			if len(subs) == 0 && src == nil {
				dst.Endpoint = nil
				continue
			} else if len(subs) == 0 {
				dst.Endpoint = src.Endpoint
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "grpc":
					if _, ok := dst.Endpoint.(*Cluster_Endpoint_Grpc); !ok {
						dst.Endpoint = &Cluster_Endpoint_Grpc{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.Endpoint.(*Cluster_Endpoint_Grpc).Grpc
						if newDst == nil {
							newDst = &Cluster_Endpoint_GRPC{}
							dst.Endpoint.(*Cluster_Endpoint_Grpc).Grpc = newDst
						}
						var newSrc *Cluster_Endpoint_GRPC
						if src != nil {
							newSrc = src.GetGrpc()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Endpoint.(*Cluster_Endpoint_Grpc).Grpc = src.GetGrpc()
						} else {
							dst.Endpoint.(*Cluster_Endpoint_Grpc).Grpc = nil
						}
					}
				case "grpc_http":
					if _, ok := dst.Endpoint.(*Cluster_Endpoint_GrpcHttp); !ok {
						dst.Endpoint = &Cluster_Endpoint_GrpcHttp{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.Endpoint.(*Cluster_Endpoint_GrpcHttp).GrpcHttp
						if newDst == nil {
							newDst = &Cluster_Endpoint_HTTP{}
							dst.Endpoint.(*Cluster_Endpoint_GrpcHttp).GrpcHttp = newDst
						}
						var newSrc *Cluster_Endpoint_HTTP
						if src != nil {
							newSrc = src.GetGrpcHttp()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Endpoint.(*Cluster_Endpoint_GrpcHttp).GrpcHttp = src.GetGrpcHttp()
						} else {
							dst.Endpoint.(*Cluster_Endpoint_GrpcHttp).GrpcHttp = nil
						}
					}
				case "http":
					if _, ok := dst.Endpoint.(*Cluster_Endpoint_Http); !ok {
						dst.Endpoint = &Cluster_Endpoint_Http{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.Endpoint.(*Cluster_Endpoint_Http).Http
						if newDst == nil {
							newDst = &Cluster_Endpoint_HTTP{}
							dst.Endpoint.(*Cluster_Endpoint_Http).Http = newDst
						}
						var newSrc *Cluster_Endpoint_HTTP
						if src != nil {
							newSrc = src.GetHttp()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Endpoint.(*Cluster_Endpoint_Http).Http = src.GetHttp()
						} else {
							dst.Endpoint.(*Cluster_Endpoint_Http).Http = nil
						}
					}
				case "mqtt":
					if _, ok := dst.Endpoint.(*Cluster_Endpoint_Mqtt); !ok {
						dst.Endpoint = &Cluster_Endpoint_Mqtt{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.Endpoint.(*Cluster_Endpoint_Mqtt).Mqtt
						if newDst == nil {
							newDst = &Cluster_Endpoint_MQTT{}
							dst.Endpoint.(*Cluster_Endpoint_Mqtt).Mqtt = newDst
						}
						var newSrc *Cluster_Endpoint_MQTT
						if src != nil {
							newSrc = src.GetMqtt()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Endpoint.(*Cluster_Endpoint_Mqtt).Mqtt = src.GetMqtt()
						} else {
							dst.Endpoint.(*Cluster_Endpoint_Mqtt).Mqtt = nil
						}
					}
				case "packet_forwarder_udp":
					if _, ok := dst.Endpoint.(*Cluster_Endpoint_PacketForwarderUdp); !ok {
						dst.Endpoint = &Cluster_Endpoint_PacketForwarderUdp{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.Endpoint.(*Cluster_Endpoint_PacketForwarderUdp).PacketForwarderUdp
						if newDst == nil {
							newDst = &Cluster_Endpoint_UDP{}
							dst.Endpoint.(*Cluster_Endpoint_PacketForwarderUdp).PacketForwarderUdp = newDst
						}
						var newSrc *Cluster_Endpoint_UDP
						if src != nil {
							newSrc = src.GetPacketForwarderUdp()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Endpoint.(*Cluster_Endpoint_PacketForwarderUdp).PacketForwarderUdp = src.GetPacketForwarderUdp()
						} else {
							dst.Endpoint.(*Cluster_Endpoint_PacketForwarderUdp).PacketForwarderUdp = nil
						}
					}
				case "backend_interfaces_http":
					if _, ok := dst.Endpoint.(*Cluster_Endpoint_BackendInterfacesHttp); !ok {
						dst.Endpoint = &Cluster_Endpoint_BackendInterfacesHttp{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.Endpoint.(*Cluster_Endpoint_BackendInterfacesHttp).BackendInterfacesHttp
						if newDst == nil {
							newDst = &Cluster_Endpoint_HTTP{}
							dst.Endpoint.(*Cluster_Endpoint_BackendInterfacesHttp).BackendInterfacesHttp = newDst
						}
						var newSrc *Cluster_Endpoint_HTTP
						if src != nil {
							newSrc = src.GetBackendInterfacesHttp()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Endpoint.(*Cluster_Endpoint_BackendInterfacesHttp).BackendInterfacesHttp = src.GetBackendInterfacesHttp()
						} else {
							dst.Endpoint.(*Cluster_Endpoint_BackendInterfacesHttp).BackendInterfacesHttp = nil
						}
					}
				case "basic_station_http":
					if _, ok := dst.Endpoint.(*Cluster_Endpoint_BasicStationHttp); !ok {
						dst.Endpoint = &Cluster_Endpoint_BasicStationHttp{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.Endpoint.(*Cluster_Endpoint_BasicStationHttp).BasicStationHttp
						if newDst == nil {
							newDst = &Cluster_Endpoint_HTTP{}
							dst.Endpoint.(*Cluster_Endpoint_BasicStationHttp).BasicStationHttp = newDst
						}
						var newSrc *Cluster_Endpoint_HTTP
						if src != nil {
							newSrc = src.GetBasicStationHttp()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Endpoint.(*Cluster_Endpoint_BasicStationHttp).BasicStationHttp = src.GetBasicStationHttp()
						} else {
							dst.Endpoint.(*Cluster_Endpoint_BasicStationHttp).BasicStationHttp = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Cluster_Endpoint_GRPC) SetFields(src *Cluster_Endpoint_GRPC, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "host":
			if len(subs) > 0 {
				return fmt.Errorf("'host' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Host = src.Host
			} else {
				var zero string
				dst.Host = zero
			}
		case "port":
			if len(subs) > 0 {
				return fmt.Errorf("'port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Port = src.Port
			} else {
				var zero uint32
				dst.Port = zero
			}
		case "tls":
			if len(subs) > 0 {
				return fmt.Errorf("'tls' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Tls = src.Tls
			} else {
				var zero bool
				dst.Tls = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Cluster_Endpoint_HTTP) SetFields(src *Cluster_Endpoint_HTTP, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "host":
			if len(subs) > 0 {
				return fmt.Errorf("'host' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Host = src.Host
			} else {
				var zero string
				dst.Host = zero
			}
		case "port":
			if len(subs) > 0 {
				return fmt.Errorf("'port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Port = src.Port
			} else {
				var zero uint32
				dst.Port = zero
			}
		case "tls":
			if len(subs) > 0 {
				return fmt.Errorf("'tls' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Tls = src.Tls
			} else {
				var zero bool
				dst.Tls = zero
			}
		case "path":
			if len(subs) > 0 {
				return fmt.Errorf("'path' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Path = src.Path
			} else {
				var zero string
				dst.Path = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Cluster_Endpoint_MQTT) SetFields(src *Cluster_Endpoint_MQTT, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "host":
			if len(subs) > 0 {
				return fmt.Errorf("'host' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Host = src.Host
			} else {
				var zero string
				dst.Host = zero
			}
		case "port":
			if len(subs) > 0 {
				return fmt.Errorf("'port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Port = src.Port
			} else {
				var zero uint32
				dst.Port = zero
			}
		case "tls":
			if len(subs) > 0 {
				return fmt.Errorf("'tls' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Tls = src.Tls
			} else {
				var zero bool
				dst.Tls = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Cluster_Endpoint_UDP) SetFields(src *Cluster_Endpoint_UDP, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "host":
			if len(subs) > 0 {
				return fmt.Errorf("'host' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Host = src.Host
			} else {
				var zero string
				dst.Host = zero
			}
		case "port":
			if len(subs) > 0 {
				return fmt.Errorf("'port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Port = src.Port
			} else {
				var zero uint32
				dst.Port = zero
			}
		case "dtls":
			if len(subs) > 0 {
				return fmt.Errorf("'dtls' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Dtls = src.Dtls
			} else {
				var zero bool
				dst.Dtls = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
