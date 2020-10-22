import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateDeceasedReceive from './components/DeceasedReceives';
import LoginPage from './components/Login';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/mainpage', WelcomePage);
    router.registerRoute('/deceasedreceive', CreateDeceasedReceive);
    router.registerRoute('/', LoginPage);
  },
});
