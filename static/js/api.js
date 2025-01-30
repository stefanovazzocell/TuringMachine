"use strict"

const API_ENDPOINT = "/api/"
const API_TIMEOUT = 10 * 1000

/**
 * Requests a new game
 * @param {string|false} id if set, it returns the game with this id
 * @param {string} difficulty the game difficulty
 * @param {number} choices sets the number of cards in the range [4, 6]
 * @returns a GameResponse or false
 */
async function fetch_game(id=false, difficulty="hard", choices=6) {
    const params = new URLSearchParams({})
    if (id) {
        params.set("id", id.replaceAll(" ", "").replaceAll("-", ""))
    } else {
        if (choices < 4 || choices > 6) return false
        params.set("choices", choices)
        if (difficulty) params.set("difficulty", difficulty)
    }
    const url = API_ENDPOINT + "game?" + params.toString()
    try {
        const response = await fetch(url, {
            signal: AbortSignal.timeout(API_TIMEOUT),
        })
        if (response.status == 200) {
            return await response.json()
        }
    } catch (error) {
        console.error(error)
        return false
    }
    return false
}

/**
 * Checks a verifier
 * @param {number} law the law to use
 * @param {string} proposal code proposed by the user
 * @returns a VerifyResponse or false
 */
async function verify_game(law, proposal) {
    const url = API_ENDPOINT + "verify?" + new URLSearchParams({
        "law": Number(law),
        "proposal": proposal,
    }).toString()
    try {
        const response = await fetch(url, {
            signal: AbortSignal.timeout(API_TIMEOUT),
        })
        if (response.status == 200) {
            return await response.json()
        }
    } catch (error) {
        console.error(error)
        return false
    }
    return false
}

/**
 * Attempts to solve a game
 * @param {number[]} criterias the number for each criteria
 * @param {number[]} verifiers the number for each verifier
 * @returns a SolverResponse or false
 */
async function solve_game(criterias, verifiers) {
    const url = API_ENDPOINT + "solve"
    try {
        const response = await fetch(url, {
            method: "POST",
            body: JSON.stringify({
                criterias: criterias,
                verifiers: verifiers,
            }),
            signal: AbortSignal.timeout(API_TIMEOUT),
        })
        if (response.status == 200) {
            return await response.json()
        }
    } catch (error) {
        console.error(error)
        return false
    }
    return false
}