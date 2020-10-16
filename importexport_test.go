// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package wallet_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/FactomProject/wallet"
)

func TestImportWithSpaces(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "wallet")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	path := fmt.Sprintf("%s%cdb", tmpDir, os.PathSeparator)

	w, err := wallet.ImportWalletFromMnemonic("yellow  yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow", path)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = w.GenerateFCTAddress()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
