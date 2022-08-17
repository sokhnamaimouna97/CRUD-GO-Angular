import { ProductDetailsComponent } 
   from './product-details/product-details.component';

import { NgModule } from '@angular/core';
import { Routes, RouterModule } 
   from '@angular/router';
import {ProductListComponent } 
  from './product-list/product-list.component';
import { UpdateProductComponent } 
   from './update-product/update-product.component';
import { CreateProductComponent } from './create-product/create-product.component';
import { SearchProductComponent } from './search-product/search-product.component';


const routes: Routes = [
  { path: '', redirectTo: 'products', pathMatch: 'full' },
  { path: 'products', component: ProductListComponent },
  { path: 'add', component: CreateProductComponent },
  { path: 'update/:id', component: UpdateProductComponent },
  { path: 'details/:id', component: ProductDetailsComponent },
  { path: 'search', component: SearchProductComponent },


];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }