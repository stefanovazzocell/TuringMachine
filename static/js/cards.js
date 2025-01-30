"use strict"

const criteria_to_laws = [2, 3, 3, 3, 2, 2, 2, 4, 4, 4, 3, 3, 3, 3, 3, 2, 4, 2,
    3, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 6, 3, 3, 3, 3, 3, 6, 9, 9, 6, 6,
    6, 6, 6, 6, 9]
const cards_scroll_sensitivity = 10

const cardsEl = document.getElementById("cards")
const cardsViewportEl = document.getElementById("cards_viewport")
const cardEls = document.querySelectorAll("#cards_viewport > .card")
const lawEls = document.querySelectorAll(".card > .laws > div")
const zoomBtnEl = document.getElementById("cards_zoom")
const verifyBtnEl = document.getElementById("cards_verify")
const verifyDialogEl = document.getElementById("verify")
const verifyProposalEl = document.getElementById("verify_proposal")
const verifyCardEl = document.getElementById("verify_card")
const verifyProposalVerifyEl = document.getElementById("verify_btn")

/**
 * Stores the cards state in the storage
 */
function store_cards_state() {
    const state = {
        laws: []
    }
    lawEls.forEach((lawEl)=>{
        if (lawEl.classList.contains("red")) {
            state.laws.push("red")
        } else if (lawEl.classList.contains("green")) {
            state.laws.push("green")
        } else {
            state.laws.push("")
        }
    })
    store_object("cards_state", state)
}

/**
 * Checks if there's a cards state
 * @returns true if a cards state is preset
 */
function has_cards_state() {
    return load("cards_state") !== null
}

/**
 * Loads the cards state from storage
 */
function restore_cards_state() {
    const state = load_object("cards_state")
    if (state === null) return
    for (let i = 0; i < state.laws.length; i++) {
        if (state.laws[i] === "") continue
        lawEls[i].classList.add(state.laws[i])
    }
}

/**
 * Clears the stored state for cards
 */
function clear_cards_state() {
    clear_storage_item("cards_state")
}

/**
 * Checks if the browser is reporting data saving or if the downlink is ~less
 * than 1.5Mbps
 * @returns true if we should use low-data mode, false otherwise
 */
function is_low_data() {
    if (navigator.connection &&
        (navigator.connection.saveData ||
            (navigator.connection.downlink &&
                navigator.connection.downlink < 1.5))) return true
    return false
}

/**
 * Checks if the screen is small (width and height <= 1080)
 * @returns true if small screen, false otherwise
 */
function is_small_screen() {
    return Math.max(window.outerHeight, window.outerWidth) <= 1080
}

/**
 * Returns the card url based on the given criteria id
 * @param {string} criteria 
 * @returns 
 */
function get_card_image_link(criteria) {
    return `images/criteria/${get_language_code()}/${criteria}.${(is_small_screen() || is_low_data()) ? "jpg" : "png"}`
}

/**
 * Returns the card url based on the given criteria id
 * @param {string} criteria 
 * @returns 
 */
function get_card_image_url(criteria) {
    return `url('${get_card_image_link(criteria)}')`
}

/**
 * This function awaits for a criteria image to load
 * @param {string} criteria 
 * @returns 
 */
async function wait_card_image_load(criteria) {
    return new Promise((resolve, reject) => {
        const image = new Image();
        image.addEventListener("load", (e)=>{
            e.preventDefault()
            resolve()
        })
        image.addEventListener("error", (e)=>{
            e.preventDefault()
            reject()
        })
        image.src = get_card_image_link(criteria)
    })
}

/**
 * Utility that waits for a series of criterias' images to load
 * @param {[]string} criterias 
 */
async function wait_all_card_image_load(criterias) {
    const promises = []
    for (let i = 0; i < criterias.length; i++) {
        promises.push(wait_card_image_load(criterias[i]))
    }
    for (let i = 0; i < criterias.length; i++) {
        await promises.pop()
    }
}

/**
 * Sets the UI cards based on given criterias.
 * There MUST be the same number of criterias and verifiers
 * @param {number[]} criterias 
 * @param {string[]} verifiers 
 */
async function set_cards(criterias, verifiers) {
    // Disable all existing cards
    cardEls.forEach((card)=>{
        card.classList.add("disabled")
    })
    // Reset the cards position
    cardsViewportEl.dataset.target = 1
    cardsViewportEl.dataset.cards = Math.min(criterias.length, cardEls.length)
    // Setup the cards
    for (let i = 0; i < criterias.length; i++) {
        if (i >= cardEls.length) break
        cardEls[i].dataset.criteria = criterias[i]
        cardEls[i].dataset.verifier = (i >= verifiers.length ? "" : verifiers[i])
        // Setup background
        cardEls[i].style.backgroundImage = get_card_image_url(criterias[i])
        // Setup verifier
        const verifier = cardEls[i].querySelector(".verifier")
        if (i >= verifiers.length || verifiers[i] == "") {
            verifier.classList.add("disabled")
        } else {
            verifier.querySelector("span").textContent = verifiers[i]
            verifier.classList.remove("disabled")
        }
        // Set the laws count
        cardEls[i].querySelector(".laws").setAttribute(
            "data-count", criteria_to_laws[criterias[i]-1])
    }
    // Reset all selections
    cardsViewportEl.querySelectorAll(".laws > div").forEach((el)=>{
        el.classList.remove("red")
        el.classList.remove("green")
    })
    // Waits for the images to load
    await wait_all_card_image_load(criterias)
    // Set the cards as visible
    for (let i = 0; i < criterias.length; i++) {
        if (i >= cardEls.length) break
        cardEls[i].classList.remove("disabled")
    }
}

/**
 * Refreshes the cards in case of language update
 */
async function update_cards_language() {
    const game = load_game()
    if (game !== null) {
        const criterias = game.criterias
        if (!criterias || criterias.length < 4) return
        for (let i = 0; i < criterias.length; i++) {
            if (i >= cardEls.length) break
            cardEls[i].style.backgroundImage = get_card_image_url(criterias[i])
        }
    }
    // Waits for the images to load
    await wait_all_card_image_load(game.criterias)
}

/**
 * Zooms in the cards
 */
function zoom_in_cards() {
    zoomBtnEl.textContent = "view all"
    cardsViewportEl.dataset.zoom = "0"
}

/**
 * Zooms out the cards
 */
function zoom_out_cards() {
    zoomBtnEl.textContent = "zoom in"
    cardsViewportEl.dataset.zoom = "1"
}

/**
 * Toggle zoom
 */
zoomBtnEl.addEventListener("mousedown", (e)=>{
    if (zoomBtnEl.classList.contains("disabled")) return
    e.preventDefault()
    if (cardsViewportEl.dataset.zoom == "1") {
        zoom_in_cards()
    } else zoom_out_cards()
})

/**
 * Checks if we're in verify mode
 * @returns true if verify mode is active, false otherwise
 */
function is_verify_mode() {
    return cardsViewportEl.dataset.verify == "1"
}

/**
 * Toggles verify mode for the cards
 */
function toggle_verify_mode() {
    if (is_verify_mode()) {
        cardsViewportEl.dataset.verify = "0"
        zoomBtnEl.classList.remove("disabled")
        verifyBtnEl.classList.remove("orange")
        return
    }
    cardsViewportEl.dataset.verify = "1"
    zoomBtnEl.classList.add("disabled")
    verifyBtnEl.classList.add("orange")
}

/**
 * Toggle verify
 */
verifyBtnEl.addEventListener("mousedown", (e)=>{
    if (verifyBtnEl.classList.contains("disabled")) return
    e.preventDefault()
    toggle_verify_mode()
})

/**
 * Register close action for verify dialog
 */
verifyDialogEl.querySelector(".close_btn").addEventListener("click", (e)=>{
    e.preventDefault()
    verifyDialogEl.close()
})

/**
 * Register handler for verify btn
 */
verifyProposalVerifyEl.addEventListener("click", async (e)=>{
    e.preventDefault()
    show_spinner()
    lock_last_round()
    const game = load_game()
    if (game === null) {
        show_alert("Error checking the verifier", 3)
        return
    }
    const request = verify_game(
        load_law(verifyDialogEl.dataset.card),
        verifyDialogEl.dataset.proposal)
    verifyBtnEl.classList.add("disabled")
    const response = await request
    if (response === false) {
        show_alert("Error checking the verifier", 3)
    } else {
        set_last_proposal_solution(
            verifyDialogEl.dataset.card,response["check"])
    }
    verifyBtnEl.classList.remove("disabled")
    verifyDialogEl.close()
    hide_spinner()
})

/**
 * Register verify pop-over
 */
cardEls.forEach((card)=>{
    card.addEventListener("click", (e)=>{
        if (!is_verify_mode()) return
        e.preventDefault()

        verifyCardEl.style.backgroundImage = get_card_image_url(card.dataset.criteria)

        const proposal = get_last_proposal()
        if (proposal === false) return

        if (count_questions_in_round(get_last_round()) >= 3) {
            show_alert("You can ask at most 3 questions per round", 3)
            return
        }

        verifyDialogEl.dataset.card = card.dataset.card
        verifyDialogEl.dataset.proposal = proposal
        verifyProposalEl.textContent = proposal
        verifyDialogEl.showModal()
    })
})

/**
 * Toggle laws
 */
lawEls.forEach((el)=>{
    el.addEventListener("contextmenu", (e)=>{
        if (is_verify_mode()) return
        e.preventDefault()
        el.classList.remove("red")
        el.classList.remove("green")
        store_cards_state()
    })
    el.addEventListener("mousedown", (e)=>{
        if (is_verify_mode()) return
        if (e.button != 0) return
        e.preventDefault()
        if (el.classList.contains("red")) {
            el.classList.remove("red")
            el.classList.add("green")
        } else if (el.classList.contains("green")) {
            el.classList.remove("green")
        } else {
            el.classList.add("red")
        }
        store_cards_state()
    })
})