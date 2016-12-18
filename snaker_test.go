package snaker

import "testing"

func TestCamelToSnake(t *testing.T) {
	var tests = []struct {
		s, exp string
	}{
		{"", ""},
		{"0", "0"},
		{"_", "_"},
		{"-X-", "-x-"},
		{"-X_", "-x_"},
		{"AReallyLongName", "a_really_long_name"},
		{"SomethingID", "something_id"},
		{"SomethingID_", "something_id_"},
		{"_SomethingID_", "_something_id_"},
		{"_Something-ID_", "_something-id_"},
		{"ACL", "acl"},
		{"GPU", "g_p_u"},
		{"zGPU", "z_g_p_u"},
		{"!GPU*", "!g_p_u*"},
		{"GpuInfo", "gpu_info"},
		{"GPUInfo", "g_p_u_info"},
		{"gpUInfo", "gp_ui_nfo"},
		{"gpUIDNfo", "gp_uid_nfo"},
		{"gpUIDnfo", "gp_uid_nfo"},
		{"HTTPWriter", "http_writer"},
		{"uHTTPWriter", "u_http_writer"},
		{"UHTTPWriter", "u_h_t_t_p_writer"},
		{"UHTTP_Writer", "u_h_t_t_p_writer"},
		{"UHTTP-Writer", "u_h_t_t_p-writer"},
		{"HTTPHTTP", "http_http"},
		{"uHTTPHTTP", "u_http_http"},
		{"uHTTPHTTPS", "u_http_https"},
		{"uHTTPHTTPS*", "u_http_https*"},
		{"uHTTPSUID*", "u_https_uid*"},
		{"UIDuuidUIDIDUUID", "uid_uuid_uid_id_uuid"},
		{"UID-uuidUIDIDUUID", "uid-uuid_uid_id_uuid"},
		{"UIDzuuidUIDIDUUID", "uid_zuuid_uid_id_uuid"},
		{"UIDzUUIDUIDidUUID", "uid_z_uuid_uid_id_uuid"},
		{"UIDzUUID-UIDidUUID", "uid_z_uuid-uid_id_uuid"},
	}

	for i, test := range tests {
		v := CamelToSnake(test.s)
		if v != test.exp {
			t.Errorf("test %d '%s' expected '%s', got: '%s'", i, test.s, test.exp, v)
		}
	}
}

func TestCamelToSnakeIdentifier(t *testing.T) {
	var tests = []struct {
		s, exp string
	}{
		{"", ""},
		{"0", ""},
		{"_", ""},
		{"-X-", "x"},
		{"-X_", "x"},
		{"AReallyLongName", "a_really_long_name"},
		{"SomethingID", "something_id"},
		{"SomethingID_", "something_id"},
		{"_SomethingID_", "something_id"},
		{"_Something-ID_", "something_id"},
		{"ACL", "acl"},
		{"GPU", "g_p_u"},
		{"zGPU", "z_g_p_u"},
		{"!GPU*", "g_p_u"},
		{"GpuInfo", "gpu_info"},
		{"GPUInfo", "g_p_u_info"},
		{"gpUInfo", "gp_ui_nfo"},
		{"gpUIDNfo", "gp_uid_nfo"},
		{"gpUIDnfo", "gp_uid_nfo"},
		{"HTTPWriter", "http_writer"},
		{"uHTTPWriter", "u_http_writer"},
		{"UHTTPWriter", "u_h_t_t_p_writer"},
		{"UHTTP_Writer", "u_h_t_t_p_writer"},
		{"UHTTP-Writer", "u_h_t_t_p_writer"},
		{"HTTPHTTP", "http_http"},
		{"uHTTPHTTP", "u_http_http"},
		{"uHTTPHTTPS", "u_http_https"},
		{"uHTTPHTTPS*", "u_http_https"},
		{"uHTTPSUID*", "u_https_uid"},
		{"UIDuuidUIDIDUUID", "uid_uuid_uid_id_uuid"},
		{"UID-uuidUIDIDUUID", "uid_uuid_uid_id_uuid"},
		{"UIDzuuidUIDIDUUID", "uid_zuuid_uid_id_uuid"},
		{"UIDzUUIDUIDidUUID", "uid_z_uuid_uid_id_uuid"},
		{"UIDzUUID-UIDidUUID", "uid_z_uuid_uid_id_uuid"},
	}

	for i, test := range tests {
		v := CamelToSnakeIdentifier(test.s)
		if v != test.exp {
			t.Errorf("test %d '%s' expected '%s', got: '%s'", i, test.s, test.exp, v)
		}
	}
}

func TestSnakeToCamel(t *testing.T) {
	var tests = []struct {
		s, exp string
	}{
		{"", ""},
		{"0", "0"},
		{"_", ""},
		{"x_", "X"},
		{"_x", "X"},
		{"_x_", "X"},
		{"a_really_long_name", "AReallyLongName"},
		{"something_id", "SomethingID"},
		{"acl", "ACL"},
		{"acl_", "ACL"},
		{"_acl", "ACL"},
		{"_acl_", "ACL"},
		{"_a_c_l_", "ACL"},
		{"gpu_info", "GpuInfo"},
		{"GPU_info", "GpuInfo"},
		{"gPU_info", "GpuInfo"},
		{"g_p_u_info", "GPUInfo"},
	}

	for i, test := range tests {
		v := SnakeToCamel(test.s)
		if v != test.exp {
			t.Errorf("test %d '%s' expected '%s', got: '%s'", i, test.s, test.exp, v)
		}
	}
}

func TestSnakeToCamelIdentifier(t *testing.T) {
	var tests = []struct {
		s, exp string
	}{
		{"", ""},
		{"_", ""},
		{"0", ""},
		{"000", ""},
		{"_000", ""},
		{"_000", ""},
		{"000_", ""},
		{"_000_", ""},
		{"___0--00_", ""},
		{"A0", "A0"},
		{"a_0", "A0"},
		{"a-0", "A0"},
		{"x_", "X"},
		{"_x", "X"},
		{"_x_", "X"},
		{"a_really_long_name", "AReallyLongName"},
		{"_a_really_long_name", "AReallyLongName"},
		{"a_really_long_name_", "AReallyLongName"},
		{"_a_really_long_name_", "AReallyLongName"},
		{"something_id", "SomethingID"},
		{"something-id", "SomethingID"},
		{"-something-id", "SomethingID"},
		{"something-id-", "SomethingID"},
		{"-something-id-", "SomethingID"},
		{"acl", "ACL"},
		{"acl_", "ACL"},
		{"_acl", "ACL"},
		{"_acl_", "ACL"},
		{"_a_c_l_", "ACL"},
		{"gpu_info", "GpuInfo"},
		{"g_p_u_info", "GPUInfo"},
	}

	for i, test := range tests {
		v := SnakeToCamelIdentifier(test.s)
		if v != test.exp {
			t.Errorf("test %d '%s' expected '%s', got: '%s'", i, test.s, test.exp, v)
		}
	}
}
