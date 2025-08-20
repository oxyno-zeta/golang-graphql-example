import { createContext } from 'react';
import { type ConfigModel, defaultConfig } from '../models/config';

export default createContext<ConfigModel>(defaultConfig);
