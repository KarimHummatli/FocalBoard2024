// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import {render, screen} from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import React from 'react'
import {Provider as ReduxProvider} from 'react-redux'

import {wrapIntl, mockStateStore} from '../../testUtils'

import {TestBlockFactory} from '../../test/testBlockFactory'

import CardDetailContentsMenu from './cardDetailContentsMenu'

//for contentRegistry
import '../content/textElement'
import '../content/imageElement'
import '../content/dividerElement'
import '../content/checkboxElement'

jest.mock('../../mutator')

const board = TestBlockFactory.createBoard()
const card = TestBlockFactory.createCard(board)
describe('components/cardDetail/cardDetailContentsMenu', () => {
    const store = mockStateStore([], {})
    beforeEach(() => {
        jest.clearAllMocks()
    })
    test('return cardDetailContentsMenu', () => {
        const {container} = render(wrapIntl(
            <ReduxProvider store={store}>
                <CardDetailContentsMenu card={card}/>
            </ReduxProvider>,
        ))
        const buttonElement = screen.getByRole('button', {name: 'menuwrapper'})
        userEvent.click(buttonElement)
        expect(container).toMatchSnapshot()
    })

    test('return cardDetailContentsMenu and add Text content', async () => {
        const {container} = render(wrapIntl(
            <ReduxProvider store={store}>
                <CardDetailContentsMenu card={card}/>
            </ReduxProvider>,
        ))
        const buttonElement = screen.getByRole('button', {name: 'menuwrapper'})
        userEvent.click(buttonElement)
        expect(container).toMatchSnapshot()
        const buttonAddText = screen.getByRole('button', {name: 'text'})
        userEvent.click(buttonAddText)
        expect(container).toMatchSnapshot()
    })
})
