// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {createAsyncThunk} from '@reduxjs/toolkit'

import {default as client} from '../octoClient'
import {Utils} from '../utils'
import {UserWorkspace} from '../user'

export const initialLoad = createAsyncThunk(
    'initialLoad',
    async () => {
        const [workspace, workspaceUsers, blocks, userWorkspaces] = await Promise.all([
            client.getWorkspace(),
            client.getWorkspaceUsers(),
            client.getAllBlocks(),
            client.getUserWorkspaces(),
        ])
        if (!workspace) {
            throw new Error('no_workspace')
        }
        return {
            workspace,
            workspaceUsers,
            blocks,
            userWorkspaces,
        }
    },
)

export const initialReadOnlyLoad = createAsyncThunk(
    'initialReadOnlyLoad',
    async (boardId: string) => {
        const blocks = client.getSubtree(boardId, 3)
        return blocks
    },
)
