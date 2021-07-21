// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
// import {IBlock} from '../blocks/block'
// import {Board, IPropertyOption, IPropertyTemplate, Board} from '../blocks/board'
// import {BoardView, BoardView} from '../blocks/boardView'
// import {Card, MutableCard} from '../blocks/card'
// import {CardFilter} from '../cardFilter'
// import {Constants} from '../constants'
// import octoClient from '../octoClient'
// import {OctoUtils} from '../octoUtils'
// import {Utils} from '../utils'
// import {IUser} from '../user'

// interface BoardTree {
//     readonly board: Board
//     readonly views: readonly BoardView[]
//     readonly cards: readonly Card[]
//     readonly cardTemplates: readonly Card[]
//     readonly allCards: readonly Card[]
//     readonly allBlocks: readonly IBlock[]

//     readonly visibleGroups: readonly Group[]
//     readonly hiddenGroups: readonly Group[]

//     readonly activeView: BoardView
//     readonly groupByProperty?: IPropertyTemplate

//     readonly rawBlocks: IBlock[]

//     getSearchText(): string | undefined
//     orderedCards(): Card[]

//     copyWithView(viewId: string): Promise<BoardTree>
//     copyWithSearchText(searchText?: string): Promise<BoardTree>
// }

// class BoardTree implements BoardTree {
//     board: Board
//     views: BoardView[] = []
//     cards: MutableCard[] = []
//     cardTemplates: MutableCard[] = []

//     visibleGroups: Group[] = []
//     hiddenGroups: Group[] = []

//     activeView!: BoardView
//     groupByProperty?: IPropertyTemplate

//     private searchText?: string
//     allCards: MutableCard[] = []
//     rawBlocks: IBlock[] = []

//     usersById: {[key: string]: IUser}

//     get allBlocks(): IBlock[] {
//         return [this.board, ...this.views, ...this.allCards, ...this.cardTemplates, ...this.rawBlocks]
//     }

//     constructor(board: Board, usersById: {[key: string]: IUser}) {
//         this.board = board
//         this.usersById = usersById
//     }

//     // Factory methods

//     static async sync(boardId: string, viewId: string, usersById: {[key: string]: IUser}): Promise<BoardTree | undefined> {
//         const rawBlocks = await octoClient.getSubtree(boardId, 3)
//         const newBoardTree = await this.buildTree(boardId, rawBlocks, usersById)
//         if (newBoardTree) {
//             newBoardTree.setActiveView(viewId)
//         }
//         return newBoardTree
//     }

//     static async incrementalUpdate(boardTree: BoardTree, updatedBlocks: IBlock[], usersById: {[key: string]: IUser}): Promise<BoardTree | undefined> {
//         const relevantBlocks = updatedBlocks.filter((block) => block.deleteAt !== 0 || block.id === boardTree.board.id || block.parentId === boardTree.board.id)
//         if (relevantBlocks.length < 1) {
//             // No change
//             return boardTree
//         }
//         const rawBlocks = OctoUtils.mergeBlocks(boardTree.allBlocks, relevantBlocks)
//         const newBoardTree = await this.buildTree(boardTree.board.id, rawBlocks, usersById)
//         newBoardTree?.setSearchText(boardTree.getSearchText())
//         if (newBoardTree && boardTree.activeView) {
//             newBoardTree.setActiveView(boardTree.activeView.id)
//         }
//         return newBoardTree
//     }

//     private static async buildTree(boardId: string, sourceBlocks: readonly IBlock[], usersById: {[key: string]: IUser}): Promise<BoardTree | undefined> {
//         const blocks = OctoUtils.hydrateBlocks(sourceBlocks)
//         const board = blocks.find((block) => block.type === 'board' && block.id === boardId) as Board
//         if (!board) {
//             return undefined
//         }
//         const boardTree = new BoardTree(board, usersById)
//         boardTree.views = blocks.filter((block) => block.type === 'view').
//             sort((a, b) => a.title.localeCompare(b.title)) as BoardView[]
//         boardTree.allCards = blocks.filter((block) => block.type === 'card' && !(block as Card).isTemplate) as MutableCard[]
//         boardTree.cardTemplates = blocks.filter((block) => block.type === 'card' && (block as Card).isTemplate).
//             sort((a, b) => a.title.localeCompare(b.title)) as MutableCard[]
//         boardTree.cards = []
//         boardTree.rawBlocks = blocks.filter((block) => block.type !== 'view' && block.type !== 'card' && block.id !== boardId)

//         boardTree.ensureMinimumSchema()
//         return boardTree
//     }

//     private ensureMinimumSchema(): boolean {
//         let didChange = false

//         // At least one select property
//         const selectProperties = this.board.cardProperties.find((o) => o.type === 'select')
//         if (!selectProperties) {
//             const newBoard = new Board(this.board)
//             newBoard.rootId = newBoard.id
//             const property: IPropertyTemplate = {
//                 id: Utils.createGuid(),
//                 name: 'Status',
//                 type: 'select',
//                 options: [],
//             }
//             newBoard.cardProperties.push(property)
//             this.board = newBoard
//             didChange = true
//         }

//         // At least one view
//         if (this.views.length < 1) {
//             const view = new BoardView()
//             view.parentId = this.board.id
//             view.rootId = this.board.rootId
//             view.groupById = this.board.cardProperties.find((o) => o.type === 'select')?.id
//             this.views.push(view)
//             didChange = true
//         }

//         if (!this.activeView) {
//             this.activeView = this.views[0]
//         }

//         return didChange
//     }

//     private setActiveView(viewId?: string): void {
//         let view: BoardView | undefined
//         if (viewId) {
//             view = this.views.find((o) => o.id === viewId)
//             if (!view) {
//                 Utils.logError(`Cannot find BoardView: ${viewId}`)
//                 view = this.views[0]
//             }
//         } else {
//             // Default to first view
//             view = this.views[0]
//         }

//         this.activeView = view!

//         // Fix missing group by (e.g. for new views)
//         if (this.activeView.viewType === 'board' && !this.activeView.groupById) {
//             this.activeView.groupById = this.board.cardProperties.find((o) => o.type === 'select')?.id
//         }

//         this.applyFilterSortAndGroup()
//     }

//     getSearchText(): string | undefined {
//         return this.searchText
//     }

//     private setSearchText(text?: string): void {
//         this.searchText = text
//         this.applyFilterSortAndGroup()
//     }

//     private applyFilterSortAndGroup(): void {
//         if (!this.activeView) {
//             Utils.assertFailure('activeView')
//             return
//         }
//         Utils.assert(this.allCards !== undefined)

//         this.cards = this.filterCards(this.allCards) as MutableCard[]
//         Utils.assert(this.cards !== undefined)
//         this.cards = this.searchFilterCards(this.cards) as MutableCard[]
//         Utils.assert(this.cards !== undefined)
//         this.cards = this.sortCards(this.cards) as MutableCard[]
//         Utils.assert(this.cards !== undefined)

//         if (this.activeView.groupById) {
//             this.setGroupByProperty(this.activeView.groupById)
//         } else {
//             Utils.assert(this.activeView.viewType !== 'board')
//         }

//         Utils.assert(this.cards !== undefined)
//     }

//     private searchFilterCards(cards: Card[]): Card[] {
//         const searchText = this.searchText?.toLocaleLowerCase()
//         if (!searchText) {
//             return cards.slice()
//         }

//         return cards.filter((card: Card) => {
//             const searchTextInCardTitle: boolean = card.title?.toLocaleLowerCase().includes(searchText)
//             if (searchTextInCardTitle) {
//                 return true
//             }

//             // Search for text in properties
//             const {board} = this
//             for (const [propertyId, propertyValue] of Object.entries(card.properties)) {
//                 // TODO: Refactor to a shared function that returns the display value of a property
//                 const propertyTemplate = board.cardProperties.find((o) => o.id === propertyId)
//                 if (propertyTemplate) {
//                     if (propertyTemplate.type === 'select') {
//                         // Look up the value of the select option
//                         const option = propertyTemplate.options.find((o) => o.id === propertyValue)
//                         if (option?.value.toLowerCase().includes(searchText)) {
//                             return true
//                         }
//                     } else if (propertyTemplate.type === 'multiSelect') {
//                         // Look up the value of the select option
//                         const options = (propertyValue as string[]).map((value) => propertyTemplate.options.find((o) => o.id === value)?.value.toLowerCase())
//                         if (options?.includes(searchText)) {
//                             return true
//                         }
//                     } else if ((propertyValue as string).toLowerCase().includes(searchText)) {
//                         return true
//                     }
//                 }
//             }

//             return false
//         })
//     }

//     private setGroupByProperty(propertyId: string) {
//         const {board} = this

//         let property = board.cardProperties.find((o) => o.id === propertyId)

//         // TODO: Handle multi-select
//         if (!property || property.type !== 'select') {
//             Utils.logError(`this.view.groupById card property not found: ${propertyId}`)
//             property = board.cardProperties.find((o) => o.type === 'select')
//             Utils.assertValue(property)
//         }
//         this.groupByProperty = property
//         this.activeView.groupById = property?.id

//         this.groupCards()
//     }

//     private groupCards() {
//         const {activeView, groupByProperty} = this
//         if (!activeView || !groupByProperty) {
//             Utils.assertFailure('groupCards')
//             return
//         }

//         const unassignedOptionIds = groupByProperty.options.
//             filter((o) => !activeView.visibleOptionIds.includes(o.id) && !activeView.hiddenOptionIds.includes(o.id)).
//             map((o) => o.id)
//         const visibleOptionIds = [...activeView.visibleOptionIds, ...unassignedOptionIds]
//         const {hiddenOptionIds} = activeView

//         // If the empty group positon is not explicitly specified, make it the first visible column
//         if (!activeView.visibleOptionIds.includes('') && !activeView.hiddenOptionIds.includes('')) {
//             visibleOptionIds.unshift('')
//         }

//         this.visibleGroups = this.groupCardsByOptions(visibleOptionIds, groupByProperty)
//         this.hiddenGroups = this.groupCardsByOptions(hiddenOptionIds, groupByProperty)
//     }

//     private groupCardsByOptions(optionIds: string[], groupByProperty: IPropertyTemplate) {
//         const groups = []
//         for (const optionId of optionIds) {
//             if (optionId) {
//                 const option = groupByProperty.options.find((o) => o.id === optionId)
//                 if (option) {
//                     const cards = this.cards.filter((o) => optionId === o.properties[groupByProperty.id])
//                     const group: Group = {
//                         option,
//                         cards,
//                     }
//                     groups.push(group)
//                 } else {
//                     Utils.logError(`groupCardsByOptions: Missing option with id: ${optionId}`)
//                 }
//             } else {
//                 // Empty group
//                 const emptyGroupCards = this.cards.filter((card) => {
//                     const groupByOptionId = card.properties[groupByProperty.id]
//                     return !groupByOptionId || !groupByProperty.options.find((option) => option.id === groupByOptionId)
//                 })
//                 const group: Group = {
//                     option: {id: '', value: `No ${groupByProperty.name}`, color: ''},
//                     cards: emptyGroupCards,
//                 }
//                 groups.push(group)
//             }
//         }

//         return groups
//     }

//     private filterCards(cards: MutableCard[]): Card[] {
//         const {board} = this
//         const filterGroup = this.activeView.filter
//         if (!filterGroup) {
//             return cards.slice()
//         }

//         return CardFilter.applyFilterGroup(filterGroup, board.cardProperties, cards)
//     }

//     private titleOrCreatedOrder(cardA: Card, cardB: Card) {
//         const aValue = cardA.title
//         const bValue = cardB.title

//         if (aValue && bValue) {
//             return aValue.localeCompare(bValue)
//         }

//         // Always put untitled cards at the bottom
//         if (aValue && !bValue) {
//             return -1
//         }
//         if (bValue && !aValue) {
//             return 1
//         }

//         // If both cards are untitled, use the create date
//         return cardA.createAt - cardB.createAt
//     }

//     private manualOrder(activeView: BoardView, cardA: Card, cardB: Card) {
//         const indexA = activeView.cardOrder.indexOf(cardA.id)
//         const indexB = activeView.cardOrder.indexOf(cardB.id)

//         if (indexA < 0 && indexB < 0) {
//             return this.titleOrCreatedOrder(cardA, cardB)
//         } else if (indexA < 0 && indexB >= 0) {
//             // If cardA's order is not defined, put it at the end
//             return 1
//         }
//         return indexA - indexB
//     }

//     private sortCards(cards: Card[]): Card[] {
//         const {board, activeView} = this
//         if (!activeView) {
//             Utils.assertFailure()
//             return cards
//         }
//         const {sortOptions} = activeView

//         if (sortOptions.length < 1) {
//             Utils.log('Manual sort')
//             return cards.sort((a, b) => this.manualOrder(activeView, a, b))
//         }

//         let sortedCards = cards
//         for (const sortOption of sortOptions) {
//             if (sortOption.propertyId === Constants.titleColumnId) {
//                 Utils.log('Sort by title')
//                 sortedCards = sortedCards.sort((a, b) => {
//                     const result = this.titleOrCreatedOrder(a, b)
//                     return sortOption.reversed ? -result : result
//                 })
//             } else {
//                 const sortPropertyId = sortOption.propertyId
//                 const template = board.cardProperties.find((o) => o.id === sortPropertyId)
//                 if (!template) {
//                     Utils.logError(`Missing template for property id: ${sortPropertyId}`)
//                     return sortedCards
//                 }
//                 Utils.log(`Sort by property: ${template?.name}`)
//                 sortedCards = sortedCards.sort((a, b) => {
//                     // Always put cards with no titles at the bottom, regardless of sort
//                     if (!a.title || !b.title) {
//                         return this.titleOrCreatedOrder(a, b)
//                     }

//                     let aValue = a.properties[sortPropertyId] || ''
//                     let bValue = b.properties[sortPropertyId] || ''

//                     if (template.type === 'createdBy') {
//                         aValue = this.usersById[a.createdBy]?.username || ''
//                         bValue = this.usersById[b.createdBy]?.username || ''
//                     } else if (template.type === 'updatedBy') {
//                         aValue = this.usersById[a.modifiedBy]?.username || ''
//                         bValue = this.usersById[b.modifiedBy]?.username || ''
//                     }

//                     let result = 0
//                     if (template.type === 'number' || template.type === 'date') {
//                         // Always put empty values at the bottom
//                         if (aValue && !bValue) {
//                             return -1
//                         }
//                         if (bValue && !aValue) {
//                             return 1
//                         }
//                         if (!aValue && !bValue) {
//                             return this.titleOrCreatedOrder(a, b)
//                         }

//                         result = Number(aValue) - Number(bValue)
//                     } else if (template.type === 'createdTime') {
//                         result = a.createAt - b.createAt
//                     } else if (template.type === 'updatedTime') {
//                         result = a.updateAt - b.updateAt
//                     } else {
//                         // Text-based sort

//                         if (aValue.length > 0 && bValue.length <= 0) {
//                             return -1
//                         }
//                         if (bValue.length > 0 && aValue.length <= 0) {
//                             return 1
//                         }
//                         if (aValue.length <= 0 && bValue.length <= 0) {
//                             return this.titleOrCreatedOrder(a, b)
//                         }

//                         if (template.type === 'select' || template.type === 'multiSelect') {
//                             aValue = template.options.find((o) => o.id === (Array.isArray(aValue) ? aValue[0] : aValue))?.value || ''
//                             bValue = template.options.find((o) => o.id === (Array.isArray(bValue) ? bValue[0] : bValue))?.value || ''
//                         }

//                         result = (aValue as string).localeCompare(bValue as string)
//                     }

//                     if (result === 0) {
//                         // In case of "ties", use the title order
//                         result = this.titleOrCreatedOrder(a, b)
//                     }

//                     return sortOption.reversed ? -result : result
//                 })
//             }
//         }

//         return sortedCards
//     }

//     orderedCards(): Card[] {
//         const cards: Card[] = []
//         for (const group of this.visibleGroups) {
//             cards.push(...group.cards)
//         }
//         for (const group of this.hiddenGroups) {
//             cards.push(...group.cards)
//         }

//         return cards
//     }

//     private async mutableCopy(): Promise<BoardTree> {
//         const x = await BoardTree.buildTree(this.board.id, this.allBlocks, this.usersById)
//         return x!
//     }

//     async copyWithView(viewId: string): Promise<BoardTree> {
//         const boardTree = await this.mutableCopy() as BoardTree
//         boardTree.setActiveView(viewId)
//         return boardTree
//     }

//     async copyWithSearchText(searchText?: string): Promise<BoardTree> {
//         const boardTree = await this.mutableCopy() as BoardTree
//         if (this.activeView) {
//             boardTree.setActiveView(this.activeView.id)
//         }
//         boardTree.setSearchText(searchText)
//         return boardTree
//     }
// }

// export {BoardTree, BoardTree}
