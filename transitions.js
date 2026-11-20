const poetryPath = new URL('poetry/', document.currentScript.src).pathname;

const isPoetryPage = url => {
    const path = new URL(url).pathname;
    return path === poetryPath.slice(0, -1) || path.startsWith(poetryPath);
};

window.addEventListener('pagereveal', event => {
    const transition = event.viewTransition;
    if (!transition) return;

    const activation = window.navigation?.activation;
    const from = activation?.from?.url;
    const to = activation?.entry?.url;
    if (!from || !to || !transition.types || window.matchMedia('(prefers-reduced-motion: reduce)').matches) {
        transition.skipTransition();
        return;
    }

    const wasPoetry = isPoetryPage(from);
    const isPoetry = isPoetryPage(to);
    if (wasPoetry === isPoetry) {
        transition.skipTransition();
        return;
    }

    transition.types.add(isPoetry ? 'poetry-enter' : 'poetry-exit');
});
