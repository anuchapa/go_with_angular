import { Routes } from '@angular/router';
import { SimpleLayout } from './layouts/simple-layout/simple-layout';

export const routes: Routes = [
    {
        path: '',
        component: SimpleLayout,
        children: [
            {
                path: '',
                loadComponent: () => import('./features/product/product-page/product-page').then(m => m.ProductPage),
            }
        ]
    }
];
