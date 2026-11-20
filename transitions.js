const poetryPath = new URL('poetry/', document.currentScript.src).pathname;
const reducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)');

const isPoetryPage = url => {
    const path = new URL(url).pathname;
    return path === poetryPath.slice(0, -1) || path.startsWith(poetryPath);
};

const poetryTransition = activation => {
    const from = activation?.from?.url;
    const to = activation?.entry?.url;
    if (!from || !to || reducedMotion.matches) return;

    const wasPoetry = isPoetryPage(from);
    const isPoetry = isPoetryPage(to);
    if (wasPoetry !== isPoetry) return isPoetry ? 'poetry-enter' : 'poetry-exit';
};

window.addEventListener('pageswap', event => {
    if (event.viewTransition && !poetryTransition(event.activation)) {
        event.viewTransition.skipTransition();
    }
});

window.addEventListener('pagereveal', event => {
    const transition = event.viewTransition;
    if (!transition) return;

    const type = poetryTransition(window.navigation?.activation);
    if (!type || !transition.types) {
        transition.skipTransition();
        return;
    }

    transition.types.add(type);
});
