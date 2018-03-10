/**
 * Copyright (c) 2018 Salesforce
 * All rights reserved.
 * Licensed under the BSD 3-Clause license.
 * For full license text, see LICENSE.txt file in the repo root
 *   or https://opensource.org/licenses/BSD-3-Clause
 */

package shuttle

import (
	"testing"
)

func TestCounter(t *testing.T) {
	counter := Counter{}
	if c := counter.Read(); c != 0 {
		t.Fatalf("counter should be 0, but was %d", c)
	}
	counter.Add(1)
	if c := counter.Read(); c != 1 {
		t.Fatalf("counter should be 1, but was %d", c)
	}
	counter.Add(2)
	if c, _ := counter.ReadAndReset(); c != 3 {
		t.Fatalf("counter should have been 3, but was %d", c)
	}
	if c := counter.Read(); c != 0 {
		t.Fatalf("counter should be have been 0 after read/reset, but was %d", c)
	}
}
