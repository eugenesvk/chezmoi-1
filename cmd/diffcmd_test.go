	"github.com/twpayne/go-vfs/v3"
	"github.com/twpayne/go-vfs/v3/vfst"
			}, func(fileSystem vfs.FS) {
					require.NoError(t, vfst.NewBuilder().Build(fileSystem, tc.extraRoot))
				require.NoError(t, newTestConfig(t, fileSystem, withStdout(&stdout)).execute(append([]string{"diff"}, tc.args...)))