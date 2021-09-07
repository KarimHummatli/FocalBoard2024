// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {createSlice, createAsyncThunk, PayloadAction} from '@reduxjs/toolkit'

import {Utils} from '../utils'

import {ClientConfig} from '../config/clientConfig'

import {default as client} from '../octoClient'

import {RootState} from './index'

export const fetchClientConfig = createAsyncThunk(
    'clientConfig/fetchClientConfig',
    async () => client.getClientConfig(),
)

const clientConfigSlice = createSlice({
    name: 'config',
    initialState: {value: {telemetry: false, telemetryid: '', enablePublicSharedBoards: false}} as {value: ClientConfig},
    reducers: {
        setClientConfig: (state, action: PayloadAction<ClientConfig>) => {
            state.value = action.payload
        },
    },
    extraReducers: (builder) => {
        builder.addCase(fetchClientConfig.fulfilled, (state, action) => {
            Utils.log('add case' + action.payload?.enablePublicSharedBoards || 'empty')
            Utils.log('add case' + action.payload?.telemetry || 'empty')
            Utils.log('add case' + action.payload?.telemetryid || 'empty')
            state.value = action.payload || {telemetry: false, telemetryid: '', enablePublicSharedBoards: false}
        })
        builder.addCase(fetchClientConfig.rejected, (state) => {
            Utils.log('rejected')
            state.value = {telemetry: false, telemetryid: '', enablePublicSharedBoards: false}
        })
    },
})

export const {setClientConfig} = clientConfigSlice.actions
export const {reducer} = clientConfigSlice

export function getClientConfig(state: RootState): ClientConfig {
    return state.clientConfig.value
}

