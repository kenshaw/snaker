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
		{"GPUInfo", "g_p_ui_nfo"},
		{"gpUInfo", "gp_ui_nfo"},
		{"gpUIDNfo", "gp_uid_nfo"},
		{"gpUIDnfo", "gp_uid_nfo"},
		{"HTTPWriter", "http_writer"},
		{"uHTTPWriter", "u_http_writer"},
		{"UHTTPWriter", "u_http_writer"},
		{"UHTTP_Writer", "u_http_writer"},
		{"UHTTP-Writer", "u_http-writer"},
		{"HTTPHTTP", "http_http"},
		{"uHTTPHTTP", "u_http_http"},
		{"uHTTPHTTPS", "u_http_https"},
		{"uHTTPHTTPS*", "u_http_https*"},
		{"uHTTPSUID*", "u_https_uid*"},
		{"UIDuuidUIDIDUUID", "uid_uuid_uid_id_uuid"},
		{"UID-uuidUIDIDUUID", "uid-uuid_uid_id_uuid"},
	}

	for i, test := range tests {
		v := CamelToSnake(test.s)
		if v != test.exp {
			t.Errorf("test %d '%s' expected %s, got: %s", i, test.s, test.exp, v)
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
		{"g_p_u_info", "GPUInfo"},
	}

	for i, test := range tests {
		v := SnakeToCamel(test.s)
		if v != test.exp {
			t.Errorf("test %d '%s' expected %s, got: %s", i, test.s, test.exp, v)
		}
	}
}
