import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    posts: {},
}

const postsSlice = createSlice({
    name: "post",
    initialState,
    reducers: {
        getPosts: (state, action) => {
            state.posts = action.payload;
        },
    },
})

export default postsSlice.reducer;

export function SetPosts(post) {
    return (dispatch, getState) => {
        dispatch(postsSlice.actions.getPosts(post));
    }
}