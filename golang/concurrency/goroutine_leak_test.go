package main

import(
    "synctest"
)

// пакет synctest из go 1.24, проверяющий утечки горутин
func ss_test(t *testing.T) {
    synctest.Test(t, func(t *testing.T) {
        <-work()
        synctest.Wait()
    })
}