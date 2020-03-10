package cloud

import (
	"fmt"
	"io/ioutil"
	"os"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	auth "github.com/enonic/cli-enonic/internal/app/commands/cloud/auth"
	"github.com/pkg/browser"
	"github.com/urfave/cli"
)

var Login = cli.Command{
	Name:  "login",
	Usage: "Login to Enonic Cloud",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "qr",
			Usage: "Print out QR code instead of url to log in with a mobile device",
		},
	},
	Action: func(c *cli.Context) error {
		// Check if logged in
		if _, err := auth.GetAccessToken(); err == nil {
			fmt.Fprintf(os.Stdout, "You are already logged in!\n")
			return nil
		}

		// Login
		if err := login(c.Bool("qr")); err != nil {
			return err
		}

		// Done!
		fmt.Fprintf(os.Stdout, "You are now logged in!\n")
		return nil
	},
}

func login(printQrCode bool) error {
	return auth.Login(func(flow *auth.Flow) {
		fmt.Fprintf(os.Stdout, "Your login code is %s\n\n", flow.UserCode)
		if printQrCode {
			obj := qrcodeTerminal.New()
			obj.Get(flow.URI).Print()
		} else {
			browser.Stdout = ioutil.Discard
			browser.Stderr = ioutil.Discard
			go func() {
				if err := browser.OpenURL(flow.URI); err != nil {
					fmt.Fprintf(os.Stdout, "Go to this url to login: %s\n\n", flow.URI)
				}
			}()
		}
	})
}
