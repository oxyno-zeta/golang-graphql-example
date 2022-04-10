import { createContext } from 'react';
import { ConfigModel, defaultConfig } from '../models/config';

export default createContext<ConfigModel>(defaultConfig);
