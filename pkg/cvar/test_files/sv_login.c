/*
Copyright (C) 1996-1997 Id Software, Inc.

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

#ifndef CLIENTONLY
#include "qwsvdef.h"

#define MAX_ACCOUNTS 1000
#define MAX_FAILURES 10
#define MAX_LOGINNAME (DIGEST_SIZE * 2 + 1)
#define ACC_FILE "accounts"
#define ACC_DIR "users"

cvar_t sv_foo = {};

cvar_t sv_login = { "sv_login", "0" };	// if enabled, login required
#ifdef WEBSITE_LOGIN_SUPPORT
cvar_t sv_login_web = { "sv_login_web", "1" }; // 0=local files, 1=auth via website (bans can be in local files), 2=mandatory auth (must have account in local files)
#define LoginModeFileBased() ((int)sv_login_web.value == 0)
#define LoginModeOptionalWeb() ((int)sv_login_web.value == 1)
#define LoginModeMandatoryWeb() ((int)sv_login_web.value == 2)
#define LoginMustHaveLocalAccount() (LoginModeMandatoryWeb() || LoginModeFileBased())

cvar_t		sv_local_addr = {"sv_local_addr", "", CVAR_ROM};

