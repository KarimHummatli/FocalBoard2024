// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {ReactElement} from 'react'
import {fireEvent, render, screen} from '@testing-library/react'

import userEvent from '@testing-library/user-event'

import '@testing-library/jest-dom'
import {IntlProvider} from 'react-intl'

import {mocked} from 'ts-jest/utils'

import mutator from '../mutator'

import {BlockIcons} from '../blockIcons'
import {TestBlockFactory} from '../test/testBlockFactory'

import BlockIconSelector from './blockIconSelector'

const wrapIntl = (children: ReactElement) => <IntlProvider locale='en'>{children}</IntlProvider>
const board = TestBlockFactory.createBoard()
const icon = BlockIcons.shared.randomIcon()

jest.mock('../mutator')
const mockedMutator = mocked(mutator, true)

describe('components/blockIconSelector', () => {
    beforeEach(() => {
        board.fields.icon = icon
        jest.clearAllMocks()
    })
    test('return an icon correctly', () => {
        render(wrapIntl(
            <BlockIconSelector
                block={board}
                size='l'
            />,
        ))
        expect(screen.queryByText(icon)).toBeTruthy()
    })
    test('return no element with no icon', () => {
        board.fields.icon = ''
        const {container} = render(wrapIntl(
            <BlockIconSelector
                block={board}
                size='l'
            />,
        ))
        expect(container).toBeEmptyDOMElement()
    })
    test('return menu on click', () => {
        render(wrapIntl(
            <BlockIconSelector
                block={board}
                size='l'
            />,
        ))
        userEvent.click(screen.getByTestId('menuwrapper'))
        expect(screen.queryByRole('button', {name: 'Random'})).not.toBeNull()
        expect(screen.queryByText('Pick icon')).not.toBeNull()
        expect(screen.queryByRole('button', {name: 'Remove icon'})).not.toBeNull()
    })
    test('return no menu in readonly', () => {
        const {container} = render(wrapIntl(
            <BlockIconSelector
                block={board}
                readonly={true}
            />,
        ))
        expect(container.querySelector('.readonly')).not.toBeEmptyDOMElement()
    })

    test('return a new icon after click on random menu', () => {
        render(wrapIntl(
            <BlockIconSelector
                block={board}
                size='l'
            />,
        ))
        userEvent.click(screen.getByTestId('menuwrapper'))
        const buttonRandom = screen.queryByRole('button', {name: 'Random'})
        expect(buttonRandom).not.toBeNull()
        userEvent.click(buttonRandom!)
        expect(mockedMutator.changeIcon).toBeCalledTimes(1)
    })

    test('return a new icon after click on EmojiPicker', async () => {
        const {container} = render(wrapIntl(
            <BlockIconSelector
                block={board}
                size='l'
            />,
        ))
        userEvent.click(screen.getByTestId('menuwrapper'))
        const menuPicker = container.querySelector('div#pick')
        expect(menuPicker).not.toBeNull()
        fireEvent.mouseEnter(menuPicker!)

        const allButtonThumbUp = await screen.findAllByRole('button', {name: /thumbsup/i})
        userEvent.click(allButtonThumbUp[0])
        expect(mockedMutator.changeIcon).toBeCalledTimes(1)
        expect(mockedMutator.changeIcon).toBeCalledWith(board.id, board.fields.icon, '👍')
    })

    test('return no icon after click on remove menu', () => {
        const {container, rerender} = render(wrapIntl(
            <BlockIconSelector
                block={board}
                size='l'
            />,
        ))
        userEvent.click(screen.getByTestId('menuwrapper'))
        const buttonRemove = screen.queryByRole('button', {name: 'Remove icon'})
        expect(buttonRemove).not.toBeNull()
        userEvent.click(buttonRemove!)
        expect(mockedMutator.changeIcon).toBeCalledTimes(1)
        expect(mockedMutator.changeIcon).toBeCalledWith(board.id, board.fields.icon, '', 'remove icon')

        //simulate reset icon
        board.fields.icon = ''

        rerender(wrapIntl(
            <BlockIconSelector
                block={board}
            />),
        )
        expect(container).toBeEmptyDOMElement()
    })
})
