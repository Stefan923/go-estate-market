import {Component, DestroyRef, inject, OnInit} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {Router, RouterLink} from "@angular/router";
import {AuthenticationService} from "../../service/authentication.service";
import {NgOptimizedImage} from "@angular/common";
import {takeUntilDestroyed} from "@angular/core/rxjs-interop";

interface LoginForm {
  email: FormControl<string>
  password: FormControl<string>
}

@Component({
  selector: 'login-component',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    NgOptimizedImage,
    RouterLink
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
})
export default class LoginComponent implements OnInit {
  protected loginForm: FormGroup<LoginForm>
  protected isSubmitting: boolean

  private destroyRef = inject(DestroyRef)

  constructor(
    private readonly router: Router,
    private readonly authenticationService: AuthenticationService,
  ) {
    this.loginForm = new FormGroup<LoginForm>({
      email: new FormControl("", {
        validators: [Validators.required],
        nonNullable: true,
      }),
      password: new FormControl("", {
        validators: [Validators.required],
        nonNullable: true,
      }),
    })
    this.isSubmitting = false
  }

  ngOnInit(): void {
  }

  submitForm(): void {
    this.isSubmitting = true

    let observable = this.authenticationService.login(this.loginForm.value as {
      email: string,
      password: string,
    })

    observable.pipe(takeUntilDestroyed(this.destroyRef)).subscribe({
      next: (response) => {
        if (response !== null && response.success) {
          this.router.navigate(["/"])
        }
        this.isSubmitting = false
      },
      error: (err) => {
        console.log(err)
        this.isSubmitting = false
      },
    })
  }
}
