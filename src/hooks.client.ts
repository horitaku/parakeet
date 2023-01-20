import {currentUser, pb} from "$lib/pocketbase";

pb.authStore.loadFromCookie(document.cookie)
pb.authStore.onChange(() => {
    //https://www.youtube.com/watch?v=AxPB3e-3yEM&t=76s

    currentUser.set(pb.authStore.model)
    document.cookie = pb.authStore.exportToCookie({httpOnly: false})
})