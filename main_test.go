package main

import (
	"testing"

	pb "github.com/dimahum/starlink-cntrl-grpc/proto"
)

func TestRequestCreation(t *testing.T) {
	// Test creating a GetStatus request
	req := &pb.Request{
		Request: &pb.Request_GetStatus{
			GetStatus: &pb.GetStatusRequest{},
		},
	}

	if req == nil {
		t.Fatal("Failed to create request")
	}

	if req.Request == nil {
		t.Fatal("Request field is nil")
	}

	switch r := req.Request.(type) {
	case *pb.Request_GetStatus:
		if r.GetStatus == nil {
			t.Fatal("GetStatus field is nil")
		}
	default:
		t.Fatalf("Unexpected request type: %T", req.Request)
	}
}

func TestResponseParsing(t *testing.T) {
	// Test creating a GetStatus response
	resp := &pb.Response{
		Response: &pb.Response_GetStatus{
			GetStatus: &pb.GetStatusResponse{
				DeviceInfo: &pb.DeviceInfo{
					Id:              "test-device-id",
					HardwareVersion: "rev1_hw",
					SoftwareVersion: "v1.0.0",
					CountryCode:     "US",
				},
				DeviceState: &pb.DeviceState{
					UptimeS: 12345.67,
				},
			},
		},
	}

	if resp == nil {
		t.Fatal("Failed to create response")
	}

	switch r := resp.Response.(type) {
	case *pb.Response_GetStatus:
		status := r.GetStatus
		if status == nil {
			t.Fatal("GetStatus response is nil")
		}

		if status.DeviceInfo == nil {
			t.Fatal("DeviceInfo is nil")
		}

		if status.DeviceInfo.Id != "test-device-id" {
			t.Errorf("Expected device ID 'test-device-id', got '%s'", status.DeviceInfo.Id)
		}

		if status.DeviceInfo.HardwareVersion != "rev1_hw" {
			t.Errorf("Expected hardware version 'rev1_hw', got '%s'", status.DeviceInfo.HardwareVersion)
		}

		if status.DeviceState == nil {
			t.Fatal("DeviceState is nil")
		}

		if status.DeviceState.UptimeS != 12345.67 {
			t.Errorf("Expected uptime 12345.67, got %f", status.DeviceState.UptimeS)
		}

	default:
		t.Fatalf("Unexpected response type: %T", resp.Response)
	}
}