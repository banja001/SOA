using Explorer.API.Startup;
using Explorer.Stakeholders.Core.Domain;

using Explorer.Stakeholders.Core.UseCases;
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Options;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.ConfigureSwagger(builder.Configuration);
const string corsPolicy = "_corsPolicy";
builder.Services.ConfigureCors(corsPolicy);
builder.Services.ConfigureAuth();

builder.Services.RegisterModules();
builder.Services.AddSingleton<PeriodicHostedService>();
builder.Services.AddHostedService(
    provider => provider.GetRequiredService<PeriodicHostedService>());

builder.Services.AddHttpClient("toursMicroservice", (client) =>
{
    var service = Environment.GetEnvironmentVariable("GO_TOUR_SERVICE_HOST") ?? "localhost";
    client.BaseAddress = new Uri($"http://{service}:8082");
});

builder.Services.AddHttpClient("encountersMicroservice", (client) =>
{
    var service = Environment.GetEnvironmentVariable("GO_ENCOUNTERS_SERVICE_HOST") ?? "localhost";
    client.BaseAddress = new Uri($"http://{service}:8090");
});

builder.Services.AddSignalR(o =>
{
    o.EnableDetailedErrors = true;
});
var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseDeveloperExceptionPage();
    app.UseSwagger();
    app.UseSwaggerUI();
}
else
{
    app.UseExceptionHandler("/error");
    app.UseHsts();
}

app.UseRouting();
app.UseCors(corsPolicy);
app.UseHttpsRedirection();
app.UseAuthorization();
//app.UseAuthentication();

app.MapHub<PublicSiteHub>("hub");
app.MapControllers();

app.Run();

// Required for automated tests
namespace Explorer.API
{
    public partial class Program { }
}