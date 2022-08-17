import { Product } from '../product';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductService } from '../product.service';
import { Observable } from 'rxjs';
import { Category } from '../category';

@Component({
  selector: 'app-create-product',
  templateUrl: './create-product.component.html',
  styleUrls: ['./create-product.component.css']
})
export class CreateProductComponent implements OnInit {

  categories: Category[]=[];
  product: Product = new Product();
  submitted = false;
  selectedProduct: any = "";
  selectedStatus: any = "";

  constructor(private productService: ProductService,
    private router: Router) { }

  ngOnInit() {
    this.reloadData();
  }

  newProduct(): void {
    this.submitted = false;
    this.product = new Product();
  }
  reloadData() {
    this.productService.getCategoriesList().subscribe(data => {
      this.categories=data
      console.log("Test data");
      
      console.log(data)
      
    })
    
  }
  onChange(changedDropdown: string) {
    if (changedDropdown === 'categorie') {
      this.selectedStatus = "";
    } else {
      this.selectedProduct = "";
    }
  }
  save() {
    this.productService
    .createProduct(this.product).subscribe(data => {
      console.log(data)
      this.product = new Product();
      this.gotoList();
    }, 
    error => console.log(error));
  }

  onSubmit() {
    this.submitted = true;
    this.save();    
  }

  gotoList() {
    this.router.navigate(['/products']);
  }
}