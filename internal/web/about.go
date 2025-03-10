// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiutil "github.com/superseriousbusiness/gotosocial/internal/api/util"
	"github.com/superseriousbusiness/gotosocial/internal/config"
	"github.com/superseriousbusiness/gotosocial/internal/gtserror"
)

const (
	aboutPath = "/about"
)

func (m *Module) aboutGETHandler(c *gin.Context) {
	instance, err := m.processor.InstanceGetV1(c.Request.Context())
	if err != nil {
		apiutil.WebErrorHandler(c, gtserror.NewErrorInternalError(err), m.processor.InstanceGetV1)
		return
	}

	c.HTML(http.StatusOK, "about.tmpl", gin.H{
		"instance":         instance,
		"languages":        config.GetInstanceLanguages().DisplayStrs(),
		"ogMeta":           ogBase(instance),
		"blocklistExposed": config.GetInstanceExposeSuspendedWeb(),
		"stylesheets": []string{
			assetsPathPrefix + "/Fork-Awesome/css/fork-awesome.min.css",
		},
		"javascript": []string{distPathPrefix + "/frontend.js"},
	})
}
