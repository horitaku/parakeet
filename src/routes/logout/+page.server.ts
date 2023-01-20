import type { Actions } from './$types'
import {redirect} from "@sveltejs/kit";

export const actions: Actions = {
    default: async ({ locals }) => {
        locals.pb.authStore.clear()
        locals.user = null
        throw redirect(303, '/')
    }
}