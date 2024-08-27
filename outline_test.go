package outline

import (
	"context"
	"testing"
)

// change the following values to match the actual values
var o = NewOutline(
	"https://31.0.0.0:33908/qwertyEXAMPLE",
	"qwertyEXAMPLEqwerty",
)

func TestOutline_AnnotatesPrometheusDataMetricsWithAutonomous(t *testing.T) {
	type args struct {
		ctx               context.Context
		asnMetricsEnabled bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:               context.Background(),
				asnMetricsEnabled: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.AnnotatesPrometheusDataMetricsWithAutonomous(tt.args.ctx, tt.args.asnMetricsEnabled); (err != nil) != tt.wantErr {
				t.Errorf("AnnotatesPrometheusDataMetricsWithAutonomous() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_ChangesTheDefaultPortForNewlyCreatedAccess(t *testing.T) {
	type args struct {
		ctx  context.Context
		port int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:  context.Background(),
				port: 8080,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.ChangesTheDefaultPortForNewlyCreatedAccess(tt.args.ctx, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("ChangesTheDefaultPortForNewlyCreatedAccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_ChangesTheHostnameForAccessKeys(t *testing.T) {
	type args struct {
		ctx      context.Context
		hostname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:      context.Background(),
				hostname: "example.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.ChangesTheHostnameForAccessKeys(tt.args.ctx, tt.args.hostname); (err != nil) != tt.wantErr {
				t.Errorf("ChangesTheHostnameForAccessKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_CreateAccessKey(t *testing.T) {
	type args struct {
		ctx context.Context
		key NewAccessKey
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				key: NewAccessKey{
					Name:  "test",
					Limit: &DataLimit{Bytes: 1000000},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := o.CreateAccessKey(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccessKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("CreateAccessKey() got = %v", got)
		})
	}
}

func TestOutline_CreateAccessKeyWithSpecificIdentifier(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		key NewAccessKey
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				id:  "3",
				key: NewAccessKey{
					Name: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := o.CreateAccessKeyWithSpecificIdentifier(tt.args.ctx, tt.args.id, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccessKeyWithSpecificIdentifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("CreateAccessKeyWithSpecificIdentifier() got = %v", got)
		})
	}
}

func TestOutline_DeleteAnAccessKey(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				id:  "3",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.DeleteAnAccessKey(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAnAccessKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_EnablesOrDisablesSharingOfMetrics(t *testing.T) {
	type args struct {
		ctx            context.Context
		metricsEnabled bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:            context.Background(),
				metricsEnabled: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.EnablesOrDisablesSharingOfMetrics(tt.args.ctx, tt.args.metricsEnabled); (err != nil) != tt.wantErr {
				t.Errorf("EnablesOrDisablesSharingOfMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_GetAnAccessKey(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				id:  "2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := o.GetAnAccessKey(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAnAccessKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Port == 0 {
				t.Errorf("GetAnAccessKey() got = %v", got)
			} else {
				t.Logf("GetAnAccessKey() got = %v", got)
			}
		})
	}
}

func TestOutline_ListAllAccessKeys(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := o.ListAllAccessKeys(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllAccessKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("ListAllAccessKeys() got = %v", got)
		})
	}
}

func TestOutline_RemoveDataLimitForTheAccessKey(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				id:  "2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.RemoveDataLimitForTheAccessKey(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RemoveDataLimitForTheAccessKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_RemovesTheAccessKeyDataLimit(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.RemovesTheAccessKeyDataLimit(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("RemovesTheAccessKeyDataLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_RenameAnAccessKey(t *testing.T) {
	type args struct {
		ctx  context.Context
		id   string
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:  context.Background(),
				id:   "2",
				name: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.RenameAnAccessKey(tt.args.ctx, tt.args.id, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("RenameAnAccessKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_RenamesTheServer(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:  context.Background(),
				name: "Moscow 2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.RenamesTheServer(tt.args.ctx, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("RenamesTheServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_ReturnsInformationAboutTheServer(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := o.ReturnsInformationAboutTheServer(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnsInformationAboutTheServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("ReturnsInformationAboutTheServer() got = %v", got)
		})
	}
}

func TestOutline_ReturnsTheDataTransferredPerAccessKey(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := o.ReturnsTheDataTransferredPerAccessKey(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnsTheDataTransferredPerAccessKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("ReturnsTheDataTransferredPerAccessKey() got = %v", got)
		})
	}
}

func TestOutline_ReturnsWhetherMetricsIsBeingShared(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := o.ReturnsWhetherMetricsIsBeingShared(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnsWhetherMetricsIsBeingShared() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("ReturnsWhetherMetricsIsBeingShared() got = %v", got)
		})
	}
}

func TestOutline_SetDataLimitForTheAccessKey(t *testing.T) {
	type args struct {
		ctx   context.Context
		id    string
		limit int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:   context.Background(),
				id:    "2",
				limit: 1000000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.SetDataLimitForTheAccessKey(tt.args.ctx, tt.args.id, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("SetDataLimitForTheAccessKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutline_SetDataTransferLimitForAllAccessKeys(t *testing.T) {
	type args struct {
		ctx   context.Context
		limit int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:   context.Background(),
				limit: 1000000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := o.SetDataTransferLimitForAllAccessKeys(tt.args.ctx, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("SetDataTransferLimitForAllAccessKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
